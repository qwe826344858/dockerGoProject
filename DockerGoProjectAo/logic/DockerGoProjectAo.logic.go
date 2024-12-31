package logic

import (
	"context"
	extProto "dockerGoProject/ExternalProto"
	proto "dockerGoProject/proto"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"log"
)

type DockerGoProjectAoLogic struct {
	// 重写 grpc proto 服务的逻辑
	proto.UnimplementedDockerGoProjectAoServer
}

func NewDockerGoProjectAoLogic() (*grpc.Server, error) {
	server := new(DockerGoProjectAoLogic)
	s := grpc.NewServer(grpc.MaxConcurrentStreams(10)) //最大连接数为10
	proto.RegisterDockerGoProjectAoServer(s, server)

	return s, nil
}

// 获取商品信息
func (logic *DockerGoProjectAoLogic) GetItemInfo(ctx context.Context, req *proto.GetItemInfoReq) (resp *proto.GetItemInfoResp, err error) {
	if req.GetItemId() < 1 {
		err = SetError("商品id入参异常")
		return
	}

	// 入参
	ClientReq := &extProto.GetItemInfoReq{
		ReqHeader: &extProto.RequestHeader{},
		ItemId: req.ItemId,
	}

	f,dockerPyClient,err :=extProto.GetDockerProjectAoClient()
	ServiceResp,err := dockerPyClient.GetItemInfo(f.Param.Ctx,ClientReq)
	if err != nil{
		log.Fatalf("GetItemInfo err:%v",err)
		return
	}

	resp = &proto.GetItemInfoResp{
		RespHeader: &proto.ResponseHeader{Errno:ServiceResp.RespHeader.Errno,Errmsg:ServiceResp.RespHeader.Errmsg},
		Id:ServiceResp.Id,
		ItemSourceName:ServiceResp.ItemSourceName,
		ItemCnName:ServiceResp.ItemCnName,
		SellOnlineCount:ServiceResp.SellOnlineCount,
		PicUrl:ServiceResp.PicUrl,
		Prices:ServiceResp.Prices,
		Currency:ServiceResp.Currency,
		Addtime:ServiceResp.Addtime,
		Modifytime:ServiceResp.Modifytime,
	}


	log.Printf("商品信息:%s", GetJsonStr(resp))
	return
}

func GetJsonStr(v any) string {
	byteList, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(byteList)
}

func SetError(strErrMsg string) error {
	return errors.New(strErrMsg)
}
