package main

import (
	"fmt"
	"net"

	pb "github.com/flyotlin/red-bean/proto"
	"github.com/flyotlin/red-bean/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const PORT = 3487

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Fatalf("failed to listen tcp port %v: [%v]", PORT, err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterImageServiceServer(s, &server.Server{})
	log.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: [%v]", err.Error())
	}
}
