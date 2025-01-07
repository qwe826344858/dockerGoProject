package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	extProto "github.com/qwe826344858/dockerGoProject/ExternalProto"
	proto "github.com/qwe826344858/dockerGoProject/proto"
	"google.golang.org/grpc"
	"log"
	"runtime"
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
	resp = new(proto.GetItemInfoResp)
	if req.GetItemId() < 1 {
		strErrorMsg := "商品id入参异常"
		resp.RespHeader = &proto.ResponseHeader{Errno:10001,Errmsg:strErrorMsg}
		return
	}

	// 入参
	ClientReq := &extProto.GetItemInfoReq{
		ReqHeader: &extProto.RequestHeader{},
		ItemId: req.GetItemId(),
	}

	f,dockerPyClient,err :=extProto.GetDockerProjectAoClient()
	defer f.CloseClient()
	ServiceResp,err := dockerPyClient.GetItemInfo(context.Background(),ClientReq)
	if err != nil{
		customPanic(fmt.Sprintf("dockerPyClient GetItemInfo itemId:%v err:%v",req.ItemId,err))
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

func customPanic(msg any) {
	// 输出错误信息
	log.Printf("fatal error: %s context:%v", msg,context.Background())
	// 打印调用栈信息
	buf := make([]byte, 1024)
	stackSize := runtime.Stack(buf, true)
	log.Printf("painc :\n%s", buf[:stackSize])
	// 结束程序
	// os.Exit(1)
}