#!/bin/bash

CACHE_DIR="$HOME/.jumpgohost/cache"
GEN_DIR="$HOME/.jumpgohost/generator.d"
if [ ! -d $CACHE_DIR ]; then
  mkdir -p $CACHE_DIR
fi

for i in $(find $GEN_DIR -name '*.sh' -type f); do [ -x $i ] && $i; done

grep --no-filename -E "^$1" $CACHE_DIR/*.findhost.tsv

