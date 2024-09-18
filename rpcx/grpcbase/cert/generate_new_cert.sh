#!/bin/bash

echo '开始生成跟证书CA...'
# 创建根证书
# 生成key
openssl genrsa -out ca.key 2048
# 生成秘钥
openssl req -new -x509 -days 7200 -key ca.key -out ca.pem

echo '开始生成服务器证书...'
# 创建server证书
#生成key
openssl ecparam -genkey -name secp384r1 -out server.key
# 生成CSR
openssl req -new -key server.key -out server.csr
# 基于CA签发证书
openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem

echo '开始生成客户端证书...'
# Client
#生成key
openssl ecparam -genkey -name secp384r1 -out client.key
#生成CSR
openssl req -new -key client.key -out client.csr
# 基于CA签发证书
openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in client.csr -out client.pem
