// agent registry
package server

import (
	"sync"
	"time"
)

type AgentInfo struct {
	ID           string
	Name         string
	Capabilities []string

	// Heartbeat-related state
	LastSeen time.Time
	Load     int32
}

type AgentRegistry struct {
	mu     sync.RWMutex
	agents map[string]AgentInfo
}

func NewAgentRegistry() *AgentRegistry {
	return &AgentRegistry{
		agents: make(map[string]AgentInfo),
	}
}

func (r *AgentRegistry) Register(agent AgentInfo) {
	r.mu.Lock()
	defer r.mu.Unlock()

	agent.LastSeen = time.Now()
	r.agents[agent.ID] = agent
}

func (r *AgentRegistry) UpdateHeartbeat(
	agentID string,
	load int32,
) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	agent, ok := r.agents[agentID]
	if !ok {
		return false
	}

	agent.LastSeen = time.Now()
	agent.Load = load
	r.agents[agentID] = agent
	return true
}

func (r *AgentRegistry) RemoveDead(timeout time.Duration) []string {
	r.mu.Lock()
	defer r.mu.Unlock()

	var dead []string
	now := time.Now()

	for id, agent := range r.agents {
		if now.Sub(agent.LastSeen) > timeout {
			dead = append(dead, id)
			delete(r.agents, id)
		}
	}
	return dead
}

func (r *AgentRegistry) List() []AgentInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := []AgentInfo{}
	for _, a := range r.agents {
		list = append(list, a)
	}
	return list
}
