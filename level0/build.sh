#!/bin/sh

#g++ -o level0 level0.cpp

cp level0.py level0
./level0 /usr/share/dict/words < short.txt > /dev/null
