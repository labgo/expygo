#!/usr/bin/env python3

try:
    with open('hugetext.txt') as fp:
        for line in fp:
            print(line.rstrip())
except IOError as ex:
    raise SystemExit(ex)
