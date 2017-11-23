#!/usr/bin/env python3

"""Somar números na linha de comando"""

import sys

total = 0.0
for arg in sys.argv[1:]:
    try:
        x = float(arg)
    except ValueError:
        print('*** número inválido:', arg)
        continue
    total += x

print('Total:', total)
