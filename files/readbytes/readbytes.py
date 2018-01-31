#!/usr/bin/env python3

try:
    with open('notsobig.dat', 'rb') as fp:
        content = fp.read()
except IOError as ex:
    raise SystemExit(ex)

print(len(content), "bytes read.")
