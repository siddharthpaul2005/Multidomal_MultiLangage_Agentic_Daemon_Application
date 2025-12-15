package proto

// AgentRequest and AgentResponse mirror the fields defined in proto/agent.proto
// and are used by the Go HTTP manager for JSON handling.
type AgentRequest struct {
	AgentName   string `json:"agent_name"`
	Action      string `json:"action"`
	PayloadJson string `json:"payload_json"`
}

type AgentResponse struct {
	AgentName  string `json:"agent_name"`
	Status     string `json:"status"`
	ResultJson string `json:"result_json"`
}
