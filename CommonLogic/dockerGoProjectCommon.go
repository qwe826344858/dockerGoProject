package CommonLogic

import (
	"context"
	extProto "dockerGoProject/ExternalProto"
	proto "dockerGoProject/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
	envcfg "dockerGoProject/DockerScript"
)

type GRpcClientParam struct {
	Ctx	 context.Context `json:"ctx"`
	Cancel context.CancelFunc `json:"cancel"`
	Conn *grpc.ClientConn `json:"conn"`
}

func GetDockerGoProjectAoClient() (param *GRpcClientParam,client proto.DockerGoProjectAoClient,err error) {

	p,errMsg  := GetServicePort("DockerGoProjectAo")
	if errMsg != ""{
		log.Fatalf("failed to GetServicePort: %v", errMsg)
		return
	}

	EnvCfg := envcfg.GetEnvConfig()
	GrpcHost := EnvCfg.GRpcConf.Host
	port := fmt.Sprintf("%s:%d",GrpcHost,p)

	// 创建带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()

	// 使用 grpc.NewClient 连接 gRPC 服务器
	// grpc.Dial和grpc.DialContext 已弃用
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	//defer conn.Close()

	// 拼接参数
	param = &GRpcClientParam{
		Ctx: ctx,
		Cancel: cancel,
		Conn: conn,
	}


	client = proto.NewDockerGoProjectAoClient(conn)
	return
}

func CloseClient(param *GRpcClientParam){
	defer param.Cancel()
	defer param.Conn.Close()

	return
}


func GetDockerProjectAoClient()(f *GRpcFactory,Client extProto.DockerProjectAoClient,err error){
	f = NewGRpcFactory()
	client,err := f.GetClient("DockerProjectAo")
	if err != nil {
		log.Fatalf("GetDockerGoProjectAoClient err:%v",err)
		return
	}

	// 类型断言为具体的客户端类型
	Client, ok := client.(extProto.DockerProjectAoClient)
	if !ok {
		err = fmt.Errorf("client is not of type DockerGoProjectAoClient")
		return
	}
	return
}
