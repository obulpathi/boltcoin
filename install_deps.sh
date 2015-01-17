#!/bin/bash

set -e

TEST_DEPS=$(go list -f '{{.Imports}} {{.TestImports}} {{.XTestImports}}' github.com/obulpathi/boltcoin/... | sed -e 's/\[//g' | sed -e 's/\]//g' | sed -e 's/C //g')
if [ "$TEST_DEPS" ]; then
  go get -race $TEST_DEPS
fi
