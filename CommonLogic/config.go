package CommonLogic

import "fmt"

type Port int64
type ServiceName string

var m = map[ServiceName]Port{
	"DockerGoProjectAo": 40001,	// go 服务
	"DockerProjectAo": 50001,	// python 服务
}

// 获取服务的分配的端口号
func GetServicePort(sName string)(port Port,strErrMsg string){
	s := ServiceName(sName)
	if _,ok := m[s]; !ok{
		strErrMsg =  fmt.Sprintf("不存在的服务名称! %s",sName)
		return
	}

	port = m[s]
	return
}


