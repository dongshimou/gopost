#!/bin/bash

git pull origin master

cd `dirname $0`

export GIN_MODE=release
export GO111MODULE=on

basepath=$(cd `dirname $0`; pwd)

#export GOPATH=$GOPATH:$basepath
#echo $GOPATH
#echo 'go path = ' $GOPATH
#cd ./src
#govendor init
#govendor add +e
#govendor get github.com/dgrijalva/jwt-go
#govendor get github.com/gin-contrib/cors
#govendor get github.com/gin-gonic/gin
#govendor get github.com/go-sql-driver/mysql
#govendor get github.com/gorilla/feeds
#govendor get github.com/jinzhu/gorm
#cd ../

go build -o ./bin/gopost

cd ./bin && ./gopost
