#!/bin/sh

FILE_PATH=${1}
FILE_PATH_MOCK=`echo ${FILE_PATH} | sed s/.go$/_mock.go/g`
PACKAGE=`grep -e '^package .*$' ${FILE_PATH} | awk '{print $2}'`

$(go env GOPATH)/bin/mockgen -source ${FILE_PATH} -package ${PACKAGE} -destination ${FILE_PATH_MOCK}