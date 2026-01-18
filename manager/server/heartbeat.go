package server

// heartbeat.go placeholder to satisfy build until heartbeat implementation is added.
// TODO: implement agent heartbeat handling.

import (
	"context"
	"log"

	pb "hyperagent/manager/proto/manager"
)

func (s *ManagerServer) Heartbeat(
	ctx context.Context,
	req *pb.HeartbeatRequest,
) (*pb.HeartbeatResponse, error) {

	ok := s.registry.UpdateHeartbeat(
		req.AgentId,
		req.Load,
	)

	if !ok {
		log.Printf("[Heartbeat] Unknown agent: %s", req.AgentId)
		return &pb.HeartbeatResponse{Ok: false}, nil
	}

	log.Printf(
		"[Heartbeat] agent=%s load=%d",
		req.AgentId,
		req.Load,
	)

	return &pb.HeartbeatResponse{Ok: true}, nil
}
