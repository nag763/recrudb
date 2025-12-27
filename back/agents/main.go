package main

import (
	"context"
	"log"
	"os"
	"recrudbagents/agents/joboffer_analyzer"
	"recrudbagents/agents/tools"

	"github.com/spf13/viper"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/geminitool"
	"google.golang.org/genai"
)

func loadConfig() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.SetDefault("MCP_ENDPOINT", "http://localhost:3000/mcp")
	viper.SetDefault("MODEL", "gemini-2.5-flash-lite")

	// If config can't be loaded, ignore
	_ = viper.ReadInConfig()
}

func main() {
	ctx := context.Background()

	loadConfig()

	model, err := gemini.NewModel(ctx, viper.GetString("MODEL"), &genai.ClientConfig{
		APIKey: viper.GetString("GOOGLE_API_KEY"),
	})

	mcpTools, err := tools.GetMcpTools([]string{"stores_offer_details"})
	if err != nil {
		log.Fatalf("couldn't load mcp tools %v", err)
	}

	jobOfferMatcher, err := joboffer_analyzer.NewJobOfferAgent(model)
	if err != nil {
		log.Fatalf("error while creating job offer agent %v", jobOfferMatcher)
	}

	orchestrator, err := llmagent.New(llmagent.Config{
		Name:        "orchestrator_agent",
		Description: "Redirects to subagents or use tools to coordinate with user prompts.",
		Model:       model,
		Instruction: `Your goal is to coordinate with other agents or use the tools available to you in order to satisfy the request.
	
		Do not perform any action on your own, your goal is to orchestrate.

		Here are what you have at your disposal:
		
		# Agents:
		
		## job_offer_agent 

		Analyzes job offers and returns key details given the user profile.
		
		If an offer is passed without further context, you should use this agent by default without asking for confirmation.

		If the first message of the session is an offer, use this agent by default.
		
		# Tools:
		## stores_offer_details
		
		Stores an offer details. This should be used only once an offer has already been analyzed by the joboffer_analyzer agent. 
		
		You don't have to wait for user confirmation to store details, but in case some informations appear to be missing, indicate it to the user.

		If the user doesn't know, search it on using the google search tool

		## google search tool

		Use it to find any missing details.
		
		`,
		Toolsets:  []tool.Toolset{mcpTools},
		Tools:     []tool.Tool{geminitool.GoogleSearch{}},
		SubAgents: []agent.Agent{jobOfferMatcher},
	})
	if err != nil {
		log.Fatalf("error while creating orchestrator agent: %v", err)
	}

	loader, err := agent.NewMultiLoader(orchestrator, jobOfferMatcher)
	if err != nil {
		log.Fatalf("error met while loading agents %v", err)
	}

	config := &launcher.Config{
		AgentLoader: loader,
	}

	l := full.NewLauncher()

	if err := l.Execute(ctx, config, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
