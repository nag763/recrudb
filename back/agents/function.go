package function

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"nag763.com/recrudbagents/agents"
)

type AgentGCFResponse struct {
	Success bool   `json:"success"`
	Error   *error `json:"error,omitempty"`
}

type AgentGCFResponseContent[T any] struct {
	AgentGCFResponse
	Content T `json:"content,omitempty"`
}

func sessionlessHttpRunner(initalizer agents.AgentInitializer) func(w http.ResponseWriter, r *http.Request) {

	runnerInitalizer := sync.OnceValues(func() (*agents.AgentRunner, error) {
		return agents.CreateAgentRunner(initalizer)
	})

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		runner, err := runnerInitalizer()
		if err != nil {
			json.NewEncoder(w).Encode(AgentGCFResponse{
				Success: false,
			})
			w.WriteHeader(http.StatusInternalServerError)
		}

		resp, err := runner.RunSessionless("")
		if err != nil {
			json.NewEncoder(w).Encode(AgentGCFResponse{
				Success: false,
			})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(AgentGCFResponseContent[string]{
			AgentGCFResponse: AgentGCFResponse{Success: true},
			Content:          resp,
		})
	}
}

func init() {
	functions.HTTP("JobOfferAnalyzer", sessionlessHttpRunner(agents.NewJobOfferAgent))
	functions.HTTP("ReachOutAgent", sessionlessHttpRunner(agents.NewReachOutAgent))
}
