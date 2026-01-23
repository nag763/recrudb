package agent

type AgentService struct {
	Runners map[string]*Runner
}

func NewAgentService(config *AgentConfig, repository *AgentRepository) (*AgentService, error) {
	runners := make(map[string]*Runner, 2)

	cvAnalyzerRunner, err := NewAgentRunner(config, repository, NewCvAnalyzer)
	if err != nil {
		return nil, err
	}
	runners["cvAnalyzer"] = cvAnalyzerRunner

	jobOfferAgent, err := NewAgentRunner(config, repository, NewJobOfferAgent)
	if err != nil {
		return nil, err
	}
	runners["jobOffer"] = jobOfferAgent

	return &AgentService{
		Runners: runners,
	}, nil
}
