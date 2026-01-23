package agent

import "google.golang.org/adk/session"

type AgentService struct {
	Runners map[string]*Runner
}

func NewAgentService(config *AgentConfig) (*AgentService, error) {
	runners := make(map[string]*Runner, 2)
	sessionService := session.InMemoryService()
	cvAnalyzerRunner, err := NewAgentRunner(config, sessionService, NewCvAnalyzer)
	if err != nil {
		return nil, err
	}
	runners["cvAnalyzer"] = cvAnalyzerRunner

	jobOfferAgent, err := NewAgentRunner(config, sessionService, NewJobOfferAgent)
	if err != nil {
		return nil, err
	}
	runners["jobOffer"] = jobOfferAgent

	return &AgentService{
		Runners: runners,
	}, nil
}
