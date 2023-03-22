#!/bin/bash

go env -w GOPROXY=https://goproxy.cn,direct

cd tiflow-6.5.1

chmod +x scripts/*.sh

go mod tidy
go mod download
go mod vendor

make build

cd -
