package CommonLogic

import (
	"context"
	envcfg "dockerGoProject/DockerScript"
	extProto "dockerGoProject/ExternalProto"
	proto "dockerGoProject/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
	"time"
)

// 客户端实例
type AoClient interface {}

func NewGRpcFactory() (f *GRpcFactory) {
	f = new(GRpcFactory)
	// 先初始化 注册表
	f.clientRegistry = make(map[ServiceName]func(*grpc.ClientConn) AoClient,0)
	// 注册客户端
	f.RegisterAllClient()

	return f
}

type GRpcFactory struct {
	mu sync.Mutex
	clientRegistry map[ServiceName]func(*grpc.ClientConn) AoClient	// 注册表
	Param *GRpcClientParam
}


func(f *GRpcFactory) GetClient(sName ServiceName) (client AoClient,err error) {

	p,errMsg  := GetServicePort(sName)
	if errMsg != ""{
		err = fmt.Errorf(errMsg)
		log.Fatalf("failed to GetServicePort: %v", errMsg)
		return
	}

	EnvCfg := envcfg.GetEnvConfig()
	GrpcHost := EnvCfg.GRpcConf.Host
	port := fmt.Sprintf("%s:%d",GrpcHost,p)

	// 创建带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	// 使用 grpc.NewClient 连接 gRPC 服务器
	// grpc.Dial和grpc.DialContext 已弃用
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	// 拼接参数
	f.Param = &GRpcClientParam{
	Ctx: ctx,
	Cancel: cancel,
	Conn: conn,
	}

	// 查找并调用注册的构造函数
	f.mu.Lock()
	constructor, exists :=f.clientRegistry[sName]
	f.mu.Unlock()
	if !exists {
		err = fmt.Errorf("unknown client type: %s", sName)
		return
	}

	// 获取客户端
	client = constructor(conn)
	return
}


func(f *GRpcFactory)CloseClient(){
	defer f.Param.Cancel()
	defer f.Param.Conn.Close()
	return
}

// 注册客户端构造函数
func(f *GRpcFactory)RegisterClient(sName ServiceName, constructor func(*grpc.ClientConn) AoClient) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.clientRegistry[sName] = constructor
	return
}

/**=======================================后续注册需要新增函数=============================================================**/

// 批量注册 也可以单个注册
func (f *GRpcFactory) RegisterAllClient()  {
	f.RegisterClient("DockerGoProjectAo",NewDockerGoProjectAoClient)
	f.RegisterClient("DockerProjectAo",NewDockerProjectAoClient)
	return
}

func NewDockerGoProjectAoClient(conn *grpc.ClientConn)(client AoClient ){
	return proto.NewDockerGoProjectAoClient(conn)
}

func NewDockerProjectAoClient(conn *grpc.ClientConn)(client AoClient ){
	return extProto.NewDockerProjectAoClient(conn)
}