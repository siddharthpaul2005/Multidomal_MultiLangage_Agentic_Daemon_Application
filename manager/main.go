package main

import (
	"log"
	"net"
	"time"

	pb "hyperagent/manager/proto/manager"
	"hyperagent/manager/server"

	"google.golang.org/grpc"
)

func main() {
	// -----------------------------
	// Core brain components
	// -----------------------------
	registry := server.NewAgentRegistry()

	// Start dead-agent reaper
	// Agents that miss heartbeats for >30s are removed
	server.StartAgentReaper(registry, 30*time.Second)

	// -----------------------------
	// Networking
	// -----------------------------
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("[Manager] failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// -----------------------------
	// gRPC Services
	// -----------------------------
	managerServer := server.NewManagerServer(registry)
	pb.RegisterManagerServer(grpcServer, managerServer)

	// -----------------------------
	// Run
	// -----------------------------
	log.Println("[Manager] Running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("[Manager] failed to serve: %v", err)
	}
}
