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
  go get -u github.com/qwe826344858/dockerGoProject
  echo "go get rpc is done" 

fi

cd proto
protoc --go_out=. --go-grpc_out=. $proto_name
echo "autogen-go finsh!  proto: $proto_name"

# ==================增加自动注册客户端函数============================#
cd "${package_name}"
# 定义要分割的字符串
input_string=$proto_name
# 定义分隔符
delimiter="."
# 使用 IFS 变量设置内部字段分隔符
IFS="$delimiter"
GrpcPbFile=""
name=""
# 遍历分割后的每个部分
for item in $input_string; do
    name="${item}"
    break
done

# 使用字符串匹配去除后缀
if [[ $input_string == *".proto" ]]; then
    GrpcPbFile="${input_string%.proto}_grpc.pb.go"
        echo "file:${GrpcPbFile}"
else
    echo "输入的字符串不包含 .proto 后缀。"
    exit 1
fi


if [ -e "${GrpcPbFile}" ]; then
  echo "文件存在 覆写代码"
  sudo chmod 777 "${GrpcPbFile}"
  sudo sed -i '15i import (\n\t"github.com/qwe826344858/dockerGoProject/CommonLogic"\n\tgrpcClient "github.com/qwe826344858/dockerGoProject/GRpcCommon"\n\t"fmt"\n)'  "${GrpcPbFile}"

 {
  echo ""
  echo ""
	echo "func Get${name}Client()(f *grpcClient.GRpcFactory,Client ${name}Client,err error){"
	echo "	var ServiceName CommonLogic.ServiceName= \"${name}\""
	echo "	f = grpcClient.NewGRpcFactory()"
	echo "	// 注册客户端"
	echo "	f.RegisterClient(ServiceName, func(conn *grpc.ClientConn) grpcClient.AoClient {return New${name}Client(conn)})"
	echo "	client,err := f.GetClient(ServiceName)"
	echo "	if err != nil {"
	echo "		return"
	echo "	}"
	echo ""
	echo "	// 类型断言为具体的客户端类型"
	echo "	Client, ok := client.(${name}Client)"
	echo "	if !ok {"
	echo "		err = fmt.Errorf(\"client is not of type ${name}Client\")"
	echo "		return"
	echo "	}"
	echo "	return"
	echo "}"

 } >> "${GrpcPbFile}"
fi

cd ..
# ==============================================================#

cd ..
go get -u ./...
echo "go get all is done"


# 生成代码后再执行
go mod tidy
echo "go mod tidy is done"


