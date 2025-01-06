package main

import (
	"context"
	"fmt"
	extProto "github.com/qwe826344858/dockerGoProject/ExternalProto"
	proto "github.com/qwe826344858/dockerGoProject/proto"
	"log"
	"runtime"
	"sync"
	"time"
)

func main(){
	// testPyService()
	// testGoService()
	BenchmarkMyFunction()
}


func BenchmarkMyFunction() {
	var wg sync.WaitGroup
	start := time.Now()
	numRequests := 10

	// TODO Fix ERR : weakly-referenced object no longer exists
	for i := 1; i < numRequests; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ret := testGoService(int64(i))
			if ret == false {
				return
			}
			// Optionally print the reply
			// fmt.Println("Reply:", reply)
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Completed %d requests in %s\n", numRequests, elapsed)
}

func testPyService(){
	f,dockerPyClient,err :=extProto.GetDockerProjectAoClient()
	defer f.CloseClient()
	if err != nil {
		log.Fatalf("GetDockerGoProjectAoClient err:%v",err)
		return
	}

	req := &extProto.GetItemInfoReq{
		ReqHeader: &extProto.RequestHeader{},
		ItemId: 2,
	}


	resp,err := dockerPyClient.GetItemInfo(context.Background(),req)
	if err != nil{
		log.Fatalf("GetItemInfo err:%v",err)
		return
	}

	fmt.Printf("testPyService resp:%v \n",resp)

	defer f.CloseClient()
}

func testGoService(i int64) bool{
	f,client,err :=proto.GetDockerGoProjectAoClient()
	id := getGoroutineID()
	fmt.Printf(" f_conn:%v\n  id:%v\n",f.Param.Conn,id)
	defer f.CloseClient()
	if err != nil {
		log.Fatalf("GetDockerGoProjectAoClient err:%v",err)
		return false
	}

	req := &proto.GetItemInfoReq{
		ReqHeader: &proto.RequestHeader{},
		ItemId: i,
	}


	resp,err := client.GetItemInfo(context.Background(),req)
	if err != nil{
		log.Fatalf("GetItemInfo err:%v",err)
		return false
	}
	fmt.Printf("testGoService resp:%v \n",resp)

	return true
}


func getGoroutineID() int {
	buf := make([]byte, 64) // 创建用于存储 goroutine 信息的缓冲区
	n := runtime.Stack(buf, true)
	for _, line := range string(buf[:n]) {
		if line == '\n' {
			break
		}
	}
	// 返回 goroutine ID
	var id int
	fmt.Sscanf(string(buf[:n]), "goroutine %d", &id)
	return id
}