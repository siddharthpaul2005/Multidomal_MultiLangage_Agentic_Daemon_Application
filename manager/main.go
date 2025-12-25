package main

import (
	pb "hyperagent/manager/proto/manager"
	"hyperagent/manager/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	//Core brain components
	registry := server.NewAgentRegistry()
	_ = registry
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	managerServer := server.NewManagerServer(registry)
	pb.RegisterManagerServer(grpcServer, managerServer)

	log.Println("[Manager] Running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
