package function

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"nag763.com/recrudbagents/agents"
)

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(NewResponseError(err))
}

func sessionlessHttpRunner(initalizer agents.AgentInitializer) func(w http.ResponseWriter, r *http.Request) {

	runnerInitalizer := sync.OnceValues(func() (*agents.AgentRunner, error) {
		return agents.CreateAgentRunner(initalizer)
	})

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		var agentGCFRequest AgentGCFRequest
		if err := json.NewDecoder(r.Body).Decode(&agentGCFRequest); err != nil {
			writeError(w, err)
			return
		}

		runner, err := runnerInitalizer()
		if err != nil {
			writeError(w, err)
			return
		}

		resp, err := runner.RunSessionless(agentGCFRequest.Content)
		if err != nil {
			writeError(w, err)
			return
		}

		json.NewEncoder(w).Encode(NewResponseSuccess(resp))
	}
}

func init() {
	functions.HTTP("JobOfferAnalyzer", sessionlessHttpRunner(agents.NewJobOfferAgent))
	functions.HTTP("ReachOutAgent", sessionlessHttpRunner(agents.NewReachOutAgent))
	functions.HTTP("CvAnalyzerAgent", sessionlessHttpRunner(agents.NewCvAnalyzer))
}
