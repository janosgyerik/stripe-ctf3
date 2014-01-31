package server

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"stripe-ctf.com/sqlcluster/log"
	"stripe-ctf.com/sqlcluster/sql"
	"stripe-ctf.com/sqlcluster/transport"
	"stripe-ctf.com/sqlcluster/util"
)

type Server struct {
	name       string
	path       string
	listen     string
	router     *mux.Router
	httpServer *http.Server
	sql        *sql.SQL
	client     *transport.Client
	cluster    *Cluster
}

// Creates a new server.
func New(path, listen string) (*Server, error) {
	cs, err := transport.Encode(listen)
	if err != nil {
		return nil, err
	}

	sqlPath := filepath.Join(path, "storage.sql")
	util.EnsureAbsent(sqlPath)

	s := &Server{
		path:    path,
		listen:  listen,
		sql:     sql.NewSQL(sqlPath),
		router:  mux.NewRouter(),
		client:  transport.NewClient(),
		cluster: NewCluster(path, cs),
	}

	return s, nil
}

// Starts the server.
func (s *Server) ListenAndServe(primary string) error {
	var err error
	// Initialize and start HTTP server.
	s.httpServer = &http.Server{
		Handler: s.router,
	}

	if primary == "" {
		s.cluster.Init()
	} else {
		cs, err := transport.Encode(primary)
		if err != nil {
			return err
		}
		s.cluster.Join(ServerAddress{
			Name:             primary,
			ConnectionString: cs,
		}, nil)
	}

	s.router.HandleFunc("/sql", s.sqlHandler).Methods("POST")
	s.router.HandleFunc("/healthcheck", s.healthcheckHandler).Methods("GET", "HEAD")
	s.router.HandleFunc("/join", s.joinHandler).Methods("POST")

	// Start Unix transport
	l, err := transport.Listen(s.listen)
	if err != nil {
		log.Fatal(err)
	}
	return s.httpServer.Serve(l)
}

func (s *Server) joinHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) delegateToPrimary(query []byte) ([]byte, error) {
	b := bytes.NewBuffer(query)
	log.Printf("delegating to primary %v: %s", s.cluster.primary, string(query))
	resp, err := s.client.SafePost(s.cluster.primary.ConnectionString, "/sql", b)
	if err != nil {
		return nil, err
	} else {
		s, _ := ioutil.ReadAll(resp)
		log.Printf("successfully delegated! <- %s", s)
		return s, nil
	}
}

// This is the only user-facing function, and accordingly the body is
// a raw string rather than JSON.
func (s *Server) sqlHandler(w http.ResponseWriter, req *http.Request) {
	query, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Couldn't read body: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	state := s.cluster.State()
	if state != "primary" {
		resp, err := s.delegateToPrimary(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(resp)
		return
	}

	log.Printf("[%s] handling query: %#v", state, string(query))

	resp, err := s.execute(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Print("[%s] Returning response to %#v: %#v", s.cluster.State(), string(query), string(resp))
	w.Write(resp)
}

func (s *Server) healthcheckHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) execute(query []byte) ([]byte, error) {
	output, err := s.sql.Execute(s.cluster.State(), string(query))

	if err != nil {
		var msg string
		if output != nil && len(output.Stderr) > 0 {
			template := `Error executing %#v (%s)

SQLite error: %s`
			msg = fmt.Sprintf(template, query, err.Error(), util.FmtOutput(output.Stderr))
		} else {
			msg = err.Error()
		}

		return nil, errors.New(msg)
	}

	formatted := fmt.Sprintf("SequenceNumber: %d\n%s",
		output.SequenceNumber, output.Stdout)
	return []byte(formatted), nil
}
