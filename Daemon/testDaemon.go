package main

import (
	"dockerGoProject/CommonLogic"
	proto "dockerGoProject/proto"
	"fmt"
	"log"
)

func main(){
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
