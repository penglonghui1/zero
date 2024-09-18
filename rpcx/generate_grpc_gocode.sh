#!/bin/bash

# 依赖项：
#
# go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
# go get -u github.com/golang/protobuf/protoc-gen-go
echo -e "\033[1;5;33m请前往【server-side/grpc_generate_tools】修改.tmpl文件\033[0m"
echo "开始使用protoc、protoc-gen-go生成golang代码..."
echo "-开始调用protoc生成golang代码,若包含HTTP接口，则生成grpc-gateway及swagger文件"
System=$(uname -s)
if [ $System == "Darwin" ]; then
    sed -i '' 's/protos\/grpc_base.proto/grpc_base.proto/g' ./protos/*.proto
else
    sed -i 's/protos\/grpc_base.proto/grpc_base.proto/g' ./protos/*.proto
fi
for file in $(ls ./protos | grep .proto); do
#  protoc -I ./protos --go_out=plugins=grpc:./protocols_rules --swagger_out=logtostderr=true:./protos --grpc-gateway_out=logtostderr=true:./protocols_rules $file
    protoc -I ./protos --go_out=plugins=grpc:./protos  $file
    echo "   - ${file} DONE!"
done


if [ $System == "Darwin" ]; then
    sed -i '' 's/grpc_base.proto/protos\/grpc_base.proto/g' ./protos/*.proto
else
    sed -i 's/grpc_base.proto/protos\/grpc_base.proto/g' ./protos/*.proto
fi
grpcgengo -r protos/ -o grpcclient
find ./ -name "*.go" | xargs gofmt -w
find ./ -name "*.go" | xargs goimports -w
echo "生成完成！"

