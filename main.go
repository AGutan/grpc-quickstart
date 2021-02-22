package main

import (
	"fmt"
	"github.com/alexgtn/grpc-quickstart/pkg/service"
	foobar "github.com/alexgtn/grpc-quickstart/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	fooBarService := service.NewService()
	foobar.RegisterFooBarServiceServer(grpcServer, fooBarService)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
