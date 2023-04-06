#!/bin/bash

chmod +x scripts/*.sh

yarn install 

go env -w GOPROXY=https://goproxy.cn,direct

go mod tidy
go mod download
go mod vendor

make build
