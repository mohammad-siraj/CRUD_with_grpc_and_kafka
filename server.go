package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mohammad-siraj/crud_kafka/crud_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//router file
func main() {
	//db initialization with sever initialization itself
	e := crud_proto.Init()
	if e != nil {
		log.Fatal(e)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8700))
	if err != nil {
		log.Fatalf("failed in initiating the server %v", err)
	}
	s := crud_proto.Server{}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	crud_proto.RegisterCarInfoServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed in initiating the server", err)
	}
}
