package function

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"nag763.com/recrudbagents/agents"
)

func sessionlessHttpRunner(initalizer agents.AgentInitializer) func(w http.ResponseWriter, r *http.Request) {

	runnerInitalizer := sync.OnceValues(func() (*agents.AgentRunner, error) {
		return agents.CreateAgentRunner(initalizer)
	})

	return func(w http.ResponseWriter, r *http.Request) {
		runner, err := runnerInitalizer()
		if err != nil {
			http.Error(w, fmt.Sprintf("runner error: %v", err), http.StatusInternalServerError)
			return
		}

		resp, err := runner.RunSessionless("")
		if err != nil {
			http.Error(w, fmt.Sprintf("agent failed to respond %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message": "Success", "agent_response1": "%v"}`, resp)
	}
}

func init() {
	functions.HTTP("JobOfferAnalyzer", sessionlessHttpRunner(agents.NewJobOfferAgent))
}
