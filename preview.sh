#!/bin/bash
set -e
subpath=`echo "$1" |tr ',' '\n' | cut -c 2- | head -1`
imgcat --width=100 $subpath
