package logic

import (
	"context"
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
	s := grpc.NewServer()
	proto.RegisterDockerGoProjectAoServer(s, server)

	return s, nil
}

// 获取商品信息
func (logic *DockerGoProjectAoLogic) GetItemInfo(ctx context.Context, req *proto.GetItemInfoReq) (resp *proto.GetItemInfoResp, err error) {
	if req.GetItemId() < 1 {
		err = SetError("商品id入参异常")
		return
	}

	resp = &proto.GetItemInfoResp{
		RespHeader: &proto.ResponseHeader{
			Errno:  0,
			Errmsg: "",
		},
		Id:             req.GetItemId(),
		ItemSourceName: "testItem",
		ItemCnName:     "测试商品",
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
