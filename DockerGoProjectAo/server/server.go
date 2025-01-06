package main

import (
	"fmt"
	"github.com/qwe826344858/dockerGoProject/CommonLogic"
	logic "github.com/qwe826344858/dockerGoProject/DockerGoProjectAo/logic"
	"log"
	"net"
)

func main() {

	p,errMsg  := CommonLogic.GetServicePort("DockerGoProjectAo")
	if errMsg != ""{
		log.Fatalf("failed to GetServicePort: %v", errMsg)
		return
	}

	port := fmt.Sprintf(":%d",p)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s,err := logic.NewDockerGoProjectAoLogic()
	if err != nil {
		log.Fatalln("Ao is failed!")
		return
	}

	fmt.Printf("gRPC server is running on port => %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
