#!/usr/bin/env python3

text = 'Nobody expects the Spanish Inquisition!!!\n'
try:
    with open('unexpected.txt', 'w') as fp:
        count = fp.write(text)
except IOError as ex:
    raise SystemExit(ex)

print(count, "bytes written.")
