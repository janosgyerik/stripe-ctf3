#!/usr/bin/env python

import sys

with open(sys.argv[1]) as f:
    content = frozenset(f.read().splitlines())

for line in sys.stdin.read().splitlines():
    res = []
    if line:
        for word in line.split(' '):
            if word.lower() in content:
                res.append(word)
            else:
                res.append('<{0}>'.format(word))
        print ' '.join(res)
    else:
        print
