#!/usr/bin/env sh

git submodule init
git submodule update

cp -r grpcgengo/templates/* ./templates
cd grpcgengo

go build -ldflags "-s -w" -o grpcgengo

cp grpcgengo $GOPATH/bin

echo "初始化完成"