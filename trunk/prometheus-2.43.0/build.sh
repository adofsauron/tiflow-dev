#!/bin/bash

go env -w GOPROXY=https://goproxy.cn,direct

npm config set registry https://registry.npm.taobao.org/    

chmod +x scripts/*.sh


cd web/ui
npm update
npm install
cd -

cd web/ui/react-app
npm update
npm install
cd -

go mod tidy
go mod download
go mod vendor

make npm_licenses

make build

# make prometheus

# make promtool
