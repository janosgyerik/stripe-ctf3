#!/usr/bin/env python

import re
import json
import subprocess

from flask import Flask, request
app = Flask(__name__)

PORT = 9090

re_chop_postfix = re.compile(r':\D.*')

index = {}

@app.route('/healthcheck')
def success():
    return '{"success": true}'

@app.route('/index')
def index():
    global work_tree, git_dir
    work_tree = request.args.get('path')
    git_dir = work_tree + '/.git'
    subprocess.check_output(['git', 'init', work_tree])
    subprocess.check_output(['git', '--git-dir', git_dir, '--work-tree', work_tree, 'add', '--all'])
    subprocess.check_output(['git', '--git-dir', git_dir, '--work-tree', work_tree, 'config', '--add', 'grep.patternType', 'fixed'])
    subprocess.check_output(['git', '--git-dir', git_dir, '--work-tree', work_tree, 'config', '--add', 'grep.lineNumber', 'true'])
    return success()

@app.route('/')
def query():
    global work_tree, git_dir
    q = request.args.get('q')
    results = re_chop_postfix.sub('', subprocess.check_output(['git', '--git-dir', git_dir, '--work-tree', work_tree, 'grep', q])).split('\n')[:-1]
    return json.dumps({
        'success': True,
        'results': results
        })

@app.route('/isIndexed')
def isIndexed():
    global git_dir
    if git_dir:
        return success()
    else:
        return '{"success": false, "error": "Nodes are not indexed"}'

if __name__ == "__main__":
    app.run(port=PORT)
