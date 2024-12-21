#!/bin/bash

proto_name=$1
package_name=$2

if [ -z "$proto_name" ] || [ -z "$package_name" ]; then
    echo "proto_name或package_name不能为空 "
    exit 1
fi

out_path="${package_name}"

if [ -d "./${package_name}" ]; then
  #如果存在对应的目录 就go mod tidy
  echo "已存在对应目录"
  cp $proto_name "./${out_path}/proto/"

  cd "${package_name}"
  
else
  echo "初始化目录"
  #不存在则初始化
  echo "create folder:$package_name"
  mkdir "${package_name}"
  mkdir "${package_name}/proto"
  cp $proto_name "./${out_path}/proto/"

  cd "${package_name}"
  go mod init "${package_name}"
  echo "go mod init is done"
  
  # 临时关闭代理
  # unset http_proxy
  # unset https_proxy 
  
  # 添加grpc依赖
  go get -u google.golang.org/grpc
  go get -u google.golang.org/protobuf/cmd/protoc-gen-go
  go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
  echo "go get rpc is done" 

fi

cd proto
protoc --go_out=. --go-grpc_out=. $proto_name
echo "autogen-go finsh!  proto: $proto_name"

cd ..
go get -u ./...
echo "go get all is done"


# 生成代码后再执行
go mod tidy
echo "go mod tidy is done"

