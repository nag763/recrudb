package joboffer_analyzer

import (
	_ "embed"
	"recrudbagents/agents/tools"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/agent/workflowagents/sequentialagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/model"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/geminitool"
)

//go:embed prompts/profile_scorer.txt
var ProfileScorerPrompt string

//go:embed prompts/offer_parser.txt
var OfferParserPrompt string

//go:embed prompts/company_lookup.txt
var CompanyLookupPrompt string

//go:embed prompts/recap_agent.txt
var RecapAgentPrompt string
var config *launcher.Config

func NewJobOfferAgent(model model.LLM) (agent.Agent, error) {

	mcpTools, err := tools.GetMcpTools([]string{"get_profile"})

	parseOfferAgent, err := llmagent.New(llmagent.Config{
		Name:        "parse_offer_agent",
		Model:       model,
		Description: "Parse a job offer.",
		Instruction: OfferParserPrompt,
		OutputKey:   "summed_up_offer",
	})
	if err != nil {
		return nil, err
	}

	lookupCompanyAgent, err := llmagent.New(llmagent.Config{
		Name:        "lookup_company",
		Model:       model,
		Description: "Look up on the company and provides a recap",
		Instruction: CompanyLookupPrompt,
		OutputKey:   "company_lookup",
		Tools:       []tool.Tool{geminitool.GoogleSearch{}},
	})
	if err != nil {
		return nil, err
	}

	offerScorerAgent, err := llmagent.New(llmagent.Config{
		Name:        "offer_scorer",
		Model:       model,
		Description: "Goal is to match my profile with the offer, on a score out of 100",
		Instruction: ProfileScorerPrompt,
		OutputKey:   "job_offer_score",
		Toolsets:    []tool.Toolset{mcpTools},
	})
	if err != nil {
		return nil, err
	}
	recapAgent, err := llmagent.New(llmagent.Config{
		Name:        "recap_agent",
		Description: "Recap informations of other agents into a few lines",
		Model:       model,
		Instruction: RecapAgentPrompt,
	})
	if err != nil {
		return nil, err
	}

	jobOfferMatcher, err := sequentialagent.New(sequentialagent.Config{
		AgentConfig: agent.Config{
			Name:        "job_offer_agent",
			Description: "Analyzes a job offer, provide a score and details.",
			SubAgents: []agent.Agent{
				parseOfferAgent,
				lookupCompanyAgent,
				offerScorerAgent,
				recapAgent,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return jobOfferMatcher, nil
}
