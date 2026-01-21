package agents

import (
	"context"
	_ "embed"
	"nag763/recrudb/api/internal/agents/config"
	"nag763/recrudb/api/internal/agents/hooks"
	"nag763/recrudb/api/internal/clients"
	"slices"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/agent/workflowagents/sequentialagent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/geminitool"
	"google.golang.org/genai"
)

//go:embed prompts/profile_scorer.txt
var profileScorerPrompt string

//go:embed prompts/offer_parser.txt
var offerParserPrompt string

//go:embed prompts/company_lookup.txt
var companyLookupPrompt string

//go:embed prompts/recap_agent.txt
var recapAgentPrompt string

//go:embed prompts/store_offer.txt
var storeOfferAgentPrompt string

func NewJobOfferAgent(ctx context.Context, model model.LLM) (agent.Agent, error) {
	offerScorerAllowedTools := []string{"get_profile"}
	offerStorerAllowedTools := []string{"stores_offer"}
	allowedTools := slices.Concat(offerScorerAllowedTools, offerStorerAllowedTools)

	mcpTools, err := clients.GetMcpTools(context.Background(), allowedTools)

	parseOfferAgent, err := llmagent.New(llmagent.Config{
		Name:        "parse_offer_agent",
		Model:       model,
		Description: "Parse a job offer.",
		Instruction: offerParserPrompt,
		OutputKey:   "summed_up_offer",
	})
	if err != nil {
		return nil, err
	}

	lookupCompanyAgent, err := llmagent.New(llmagent.Config{
		Name:        "lookup_company",
		Model:       model,
		Description: "Look up on the company and provides a recap",
		Instruction: companyLookupPrompt,
		OutputKey:   "company_lookup",
		Tools:       []tool.Tool{geminitool.GoogleSearch{}},
	})
	if err != nil {
		return nil, err
	}

	offerScorerAgent, err := llmagent.New(llmagent.Config{
		Name:                "offer_scorer",
		Model:               model,
		Description:         "Goal is to match my profile with the offer, on a score out of 100",
		Instruction:         profileScorerPrompt,
		OutputKey:           "job_offer_score",
		Toolsets:            []tool.Toolset{mcpTools},
		BeforeToolCallbacks: []llmagent.BeforeToolCallback{hooks.InjectMetadata},
		GenerateContentConfig: &genai.GenerateContentConfig{
			ToolConfig: &genai.ToolConfig{
				FunctionCallingConfig: config.RestrictConfigTo(offerScorerAllowedTools),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	recapAgent, err := llmagent.New(llmagent.Config{
		Name:        "recap_agent",
		Description: "Recap informations of other agents into a few lines",
		Model:       model,
		Instruction: recapAgentPrompt,
		OutputKey:   "recap",
	})
	if err != nil {
		return nil, err
	}

	storeOfferAgent, err := llmagent.New(llmagent.Config{
		Name:        "store_offer_agent",
		Description: "Stores the offer details",
		Model:       model,
		Instruction: storeOfferAgentPrompt,
		Toolsets:    []tool.Toolset{mcpTools},
		GenerateContentConfig: &genai.GenerateContentConfig{
			ToolConfig: &genai.ToolConfig{
				FunctionCallingConfig: config.RestrictConfigTo(offerStorerAllowedTools),
			},
		},
	})

	jobOfferMatcher, err := sequentialagent.New(sequentialagent.Config{
		AgentConfig: agent.Config{
			Name:        "job_offer_agent",
			Description: "Analyzes a job offer, provide a score and details.",
			SubAgents: []agent.Agent{
				parseOfferAgent,
				lookupCompanyAgent,
				offerScorerAgent,
				recapAgent,
				storeOfferAgent,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return jobOfferMatcher, nil
}
