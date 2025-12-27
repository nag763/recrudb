package main

import (
	"context"
	_ "embed"
	"log"
	"os"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/spf13/viper"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/agent/workflowagents/sequentialagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/mcptoolset"
	"google.golang.org/genai"
)

//go:embed prompts/profile_scorer.txt
var ProfileScorerPrompt string

//go:embed prompts/offer_parser.txt
var OfferParserPrompt string

//go:embed prompts/company_lookup.txt
var CompanyLookupPrompt string

//go:embed prompts/recap_agent.txt
var RecapAgentPrompt string

var ctx context.Context
var jobOfferMatcher agent.Agent
var config *launcher.Config

func init() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.SetDefault("MCP_ENDPOINT", "http://localhost:3000/mcp")
	viper.SetDefault("MODEL", "gemini-2.5-flash-lite")
	// If config can't be loaded, ignore
	_ = viper.ReadInConfig()

	ctx = context.Background()

	toolset, err := mcptoolset.New(mcptoolset.Config{
		Transport: &mcp.StreamableClientTransport{
			Endpoint: viper.GetString("MCP_ENDPOINT"),
		},
	})
	if err != nil {
		log.Fatalf("getting tools %v", err)
	}

	model, err := gemini.NewModel(ctx, viper.GetString("MODEL"), &genai.ClientConfig{
		APIKey: viper.GetString("GOOGLE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("creating model: %v", err)
	}

	parseOfferAgent, err := llmagent.New(llmagent.Config{
		Name:        "parse_offer_agent",
		Model:       model,
		Description: "Parse a job offer.",
		Instruction: OfferParserPrompt,
		OutputKey:   "summed_up_offer",
		// Tools:       []tool.Tool{},
	})
	if err != nil {
		log.Fatalf("creating parseOfferAgent: %v", err)
	}

	lookupCompanyAgent, err := llmagent.New(llmagent.Config{
		Name:        "lookup_company",
		Model:       model,
		Description: "Look up on the company and provides a recap",
		Instruction: CompanyLookupPrompt,
		OutputKey:   "company_lookup",
	})
	if err != nil {
		log.Fatalf("creating lookupCompanyAgent: %v", err)
	}

	offerScorerAgent, err := llmagent.New(llmagent.Config{
		Name:        "offer_scorer",
		Model:       model,
		Description: "Goal is to match my profile with the offer, on a score out of 100",
		Instruction: ProfileScorerPrompt,
		OutputKey:   "job_offer_score",
		Toolsets:    []tool.Toolset{toolset},
	})
	if err != nil {
		log.Fatalf("creating offerScorerAgent: %v", err)
	}

	recapAgent, err := llmagent.New(llmagent.Config{
		Name:        "recap_agent",
		Description: "Recap informations of other agents into a few lines",
		Model:       model,
		Instruction: RecapAgentPrompt,
	})
	if err != nil {
		log.Fatalf("synth agent %v", err)
	}

	jobOfferMatcher, err = sequentialagent.New(sequentialagent.Config{
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
		log.Fatalf("creating jobOfferMatcher: %v", err)
	}

	config = &launcher.Config{
		AgentLoader: agent.NewSingleLoader(jobOfferMatcher),
	}
}

func main() {
	l := full.NewLauncher()

	if err := l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
