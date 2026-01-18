// routing logic
package server

func (r *AgentRegistry) FindByCapability(cap string) *AgentInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, agent := range r.agents {
		for _, c := range agent.Capabilities {
			if c == cap {
				a := agent // copy to avoid pointer alias bug
				return &a
			}
		}
	}
	return nil
}
