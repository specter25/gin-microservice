#!/usr/bin/env bash 
set -xe 

# install package and dependencies
go get github.com/gin-gonic/gin

go get gopkg.in/go-playground/validator/v10


# build command
go build -o bin/gin-microservice server.go