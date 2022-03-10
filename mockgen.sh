#!/bin/sh

set -e

files=`egrep "^type .+ interface" -lr internal | grep -v _mock.go`

for f in ${files}; do
    mockgen -source ${f} -destination `dirname ${f}`/mock/`basename ${f}`
done