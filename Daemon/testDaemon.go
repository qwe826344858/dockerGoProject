package main

import (
	"dockerGoProject/CommonLogic"
	proto "dockerGoProject/proto"
	"fmt"
	"log"
)

func main(){
	test1()
	test2()
}


func test1(){
	f := CommonLogic.NewGRpcFactory()
	client,err := f.GetClient("DockerGoProjectAo")
	if err != nil {
		log.Fatalf("GetDockerGoProjectAoClient err:%v",err)
		return
	}

	req := &proto.GetItemInfoReq{
		ReqHeader: &proto.RequestHeader{},
		ItemId: 2,
	}

	// 类型断言为具体的客户端类型
	dockerGoClient, ok := client.(proto.DockerGoProjectAoClient)
	if !ok {
		log.Fatalf("client is not of type DockerGoProjectAoClient")
		return
	}

	resp,err := dockerGoClient.GetItemInfo(f.Param.Ctx,req)
	if err != nil{
		log.Fatalf("GetItemInfo err:%v",err)
		return
	}

	fmt.Sprintf("resp:%v",resp)

	defer f.CloseClient()
}


func test2(){
	param,client,err := CommonLogic.GetDockerGoProjectAoClient()
	if err != nil {
		log.Fatalf("GetDockerGoProjectAoClient err:%v",err)
		return
	}

	req := &proto.GetItemInfoReq{
		ReqHeader: &proto.RequestHeader{},
		ItemId: 2,
	}

	resp,err := client.GetItemInfo(param.Ctx,req)
	if err != nil{
		log.Fatalf("GetItemInfo err:%v",err)
		return
	}

	fmt.Sprintf("resp:%v",resp)

	defer CommonLogic.CloseClient(param)
}