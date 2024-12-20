package CommonLogic

import "fmt"

type Port int64
type ServiceName string

var m = map[ServiceName]Port{
"DockerGoProjectAo": 40001,
}

// 获取服务的分配的端口号
func GetServicePort(sName ServiceName)(port Port,strErrMsg string){
	if _,ok := m[sName]; !ok{
		strErrMsg =  fmt.Sprintf("不存在的服务名称! %s",sName)
		return
	}

	port = m[sName]
	return
}


