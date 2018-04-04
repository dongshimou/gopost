#!/bin/bash


cd `dirname $0`
cd ../

export GIN_MODE=release

basepath=$(cd `dirname $0`; pwd)

echo $GOPATH
echo $basepath

export GOPATH=$basepath:$GOPATH

echo $GOPATH

rm -rf ./pkg

cd ./src/gopost

echo 'go path = ' $GOPATH

rm -rf ./vendor

govendor init

govendor add +e

export GOPATH=$basepath

echo $GOPATH

go build -o ../../bin/gopost

cd ../../bin && ./gopost
