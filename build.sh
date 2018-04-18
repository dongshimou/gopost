#!/bin/bash


cd `dirname $0`

export GIN_MODE=release

basepath=$(cd `dirname $0`; pwd)

export GOPATH=$basepath:$GOPATH

echo $GOPATH

echo 'go path = ' $GOPATH

govendor init

govendor add +e

govendor get github.com/dgrijalva/jwt-go

govendor get github.com/gin-contrib/cors

govendor get github.com/gin-gonic/gin

govendor get github.com/go-sql-driver/mysql

govendor get github.com/gorilla/feeds

govendor get github.com/jinzhu/gorm

go build -o ../bin/gopost

cd ../bin && ./gopost
