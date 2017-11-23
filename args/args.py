#!/usr/bin/env python3

"""Exibir argumentos da linha de comando"""

import sys

for i, arg in enumerate(sys.argv):
    print(f'{i}: {arg}')
