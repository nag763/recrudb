package function

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/genai"
	"nag763.com/recrudbagents/agents"
	"nag763.com/recrudbagents/utils"
)

func init() {
	functions.HTTP("HelloWorld", HelloWorld)
}

var (
	getModel = sync.OnceValues(func() (model.LLM, error) {
		return gemini.NewModel(context.Background(), utils.GetEnvOrDefault("MODEL", "gemini-2.5-flash-lite"), &genai.ClientConfig{
			Backend: genai.BackendVertexAI,
			Project: utils.GetEnvOrDefault("PROJECT_NAME", "recrudb"),
		})
	})

	getJobOfferAgent = sync.OnceValues(func() (agent.Agent, error) {
		m, err := getModel()
		if err != nil {
			return nil, fmt.Errorf("model init: %w", err)
		}
		return agents.NewReachOutAgent(m)
	})
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jobOfferMatcher, err := getJobOfferAgent()
	if err != nil {
		http.Error(w, fmt.Sprintf("Configuration Error: %v", err), http.StatusInternalServerError)
		return
	}

	resp, err := utils.CallAgent(ctx, jobOfferMatcher, "A job offer for C#, .NET at Google")
	if err != nil {
		http.Error(w, fmt.Sprintf("Agent failed to respond %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Success", "agent_response": "%v"}`, resp)
}
