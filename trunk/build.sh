#!/bin/bash

go env -w GOPROXY=https://goproxy.cn,direct

cd tiflow-6.5.1

go mod download
go mod vendor

make build

cd -