#!/bin/bash

go env -w GOPROXY=https://goproxy.cn,direct


chmod +x scripts/*.sh

yarn install --pure-lockfile

go mod tidy
go mod download
go mod vendor

make build
