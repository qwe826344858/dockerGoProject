package main

import (
	extProto "dockerGoProject/ExternalProto"
	proto "dockerGoProject/proto"
	"fmt"
	"log"
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
	numRequests := 1000

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ret := testGoService()
			if ret == false {
				fmt.Println("Error:")
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


	resp,err := dockerPyClient.GetItemInfo(f.Param.Ctx,req)
	if err != nil{
		log.Fatalf("GetItemInfo err:%v",err)
		return
	}

	fmt.Printf("testPyService resp:%v \n",resp)

	defer f.CloseClient()
}

func testGoService() bool{
	f,client,err :=proto.GetDockerGoProjectAoClient()
	defer f.CloseClient()
	if err != nil {
		log.Fatalf("GetDockerGoProjectAoClient err:%v",err)
		return false
	}

	req := &proto.GetItemInfoReq{
		ReqHeader: &proto.RequestHeader{},
		ItemId: 2,
	}


	resp,err := client.GetItemInfo(f.Param.Ctx,req)
	if err != nil{
		log.Fatalf("GetItemInfo err:%v",err)
		return false
	}
	fmt.Printf("testPyService resp:%v \n",resp)

	return false
}