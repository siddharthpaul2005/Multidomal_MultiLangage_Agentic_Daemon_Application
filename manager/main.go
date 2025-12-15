package main

import (
	"encoding/json"
	"log"
	"net/http"

	agentpb "hyperagent/proto"
)

func runAgentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var req agentpb.AgentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	log.Printf("Received AgentRequest: %+v\n", req)

	resp := &agentpb.AgentResponse{
		AgentName:  req.AgentName,
		Status:     "ok",
		ResultJson: `{"message":"dummy response"}`,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}

func main() {
	http.HandleFunc("/run-agent", runAgentHandler)

	addr := ":8080"
	log.Printf("Starting manager HTTP server on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}