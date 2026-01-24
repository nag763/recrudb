package agent

import "google.golang.org/adk/session"

type AgentService struct {
	Runners map[string]*Runner
}

func NewAgentService(config *AgentConfig) (*AgentService, error) {
	sessionService := session.InMemoryService()

	definitions := map[string]AgentInitializer{
		"/cvAnalyzer": NewCvAnalyzer,
		"/jobOffer":   NewJobOfferAgent,
	}

	runners := make(map[string]*Runner, len(definitions))
	for path, init := range definitions {
		runner, err := NewAgentRunner(config, sessionService, init)
		if err != nil {
			return nil, err
		}
		runners[path] = runner
	}

	return &AgentService{
		Runners: runners,
	}, nil
}
