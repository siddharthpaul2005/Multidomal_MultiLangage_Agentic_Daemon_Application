// gRPC gateway
package server

import (
	"context"
	"log"

	pb "hyperagent/manager/proto/manager"
)

type ManagerServer struct {
	pb.UnimplementedManagerServer
	registry *AgentRegistry
}

func NewManagerServer(registry *AgentRegistry) *ManagerServer {
	return &ManagerServer{
		registry: registry,
	}
}

func (s *ManagerServer) RegisterAgent(
	ctx context.Context,
	req *pb.RegisterAgentRequest,
) (*pb.RegisterAgentResponse, error) {

	agent := AgentInfo{
		ID:           req.AgentName, // TODO: replace with UUID
		Name:         req.AgentName,
		Capabilities: req.Capabilities,
	}

	s.registry.Register(agent)

	log.Printf(
		"[Manager] Registered agent: %s (%v)",
		agent.Name,
		agent.Capabilities,
	)

	return &pb.RegisterAgentResponse{
		Accepted: true,
		AgentId:  agent.ID,
	}, nil
}
