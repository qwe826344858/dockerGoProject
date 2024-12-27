package main

import (
	"dockerGoProject/CommonLogic"
	extProto "dockerGoProject/ExternalProto"
	proto "dockerGoProject/proto"
	"fmt"
	"log"
)

func main(){
	test1()
	test2()
	testPyService()
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

	fmt.Printf("test1 resp:%v \n",resp)

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

	fmt.Printf("test2 resp:%v \n",resp)

	defer CommonLogic.CloseClient(param)
}


func testPyService(){
	f := CommonLogic.NewGRpcFactory()
	client,err := f.GetClient("DockerProjectAo")
	if err != nil {
		log.Fatalf("GetDockerGoProjectAoClient err:%v",err)
		return
	}

	req := &extProto.GetItemInfoReq{
		ReqHeader: &extProto.RequestHeader{},
		ItemId: 2,
	}

	// 类型断言为具体的客户端类型
	dockerPyClient, ok := client.(extProto.DockerProjectAoClient)
	if !ok {
		log.Fatalf("client is not of type DockerGoProjectAoClient")
		return
	}

	resp,err := dockerPyClient.GetItemInfo(f.Param.Ctx,req)
	if err != nil{
		log.Fatalf("GetItemInfo err:%v",err)
		return
	}

	fmt.Printf("testPyService resp:%v \n",resp)

	defer f.CloseClient()
}