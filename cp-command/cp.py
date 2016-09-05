#!/usr/bin/python
import sys
import shutil

if len(sys.argv) == 3:
    input = sys.argv[1]
    output = sys.argv[2]
    shutil.comp2(input, output)
