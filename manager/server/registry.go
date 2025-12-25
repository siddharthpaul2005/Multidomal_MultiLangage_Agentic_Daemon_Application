// agent registry
package server

import "sync"

type AgentInfo struct {
	ID           string
	Name         string
	Capabilities []string
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
	r.agents[agent.ID] = agent
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
