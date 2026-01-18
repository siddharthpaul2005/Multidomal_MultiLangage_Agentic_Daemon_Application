package server

import (
	"log"
	"time"
)

// StartAgentReaper periodically removes agents
// that have not sent heartbeats within `timeout`
func StartAgentReaper(
	registry *AgentRegistry,
	timeout time.Duration,
) {
	ticker := time.NewTicker(timeout / 2)

	go func() {
		for range ticker.C {
			dead := registry.RemoveDead(timeout)
			for _, id := range dead {
				log.Printf("[Reaper] Agent %s declared DEAD", id)
			}
		}
	}()
}
