#!/bin/bash

proto_name=$1
project_name=$2

if [ -z "$proto_name" ] || [ -z "$project_name" ]; then
    echo "proto_name或project_name不能为空 "
    exit 1
fi



out_path="${project_name}/proto"

if [ -d "./${project_name}" ]; then
  echo "1"
else
  echo "create folder:$project_name"
  mkdir "${project_name}"


fi

if [ -d "./${project_name}/proto" ]; then
   echo "2"
else
   echo "create proto"
   mkdir "./${project_name}/proto"

fi


python3 -m  grpc_tools.protoc -I ./ --python_out="./${out_path}" --pyi_out="./${out_path}" --grpc_python_out="./${out_path}" $proto_name




