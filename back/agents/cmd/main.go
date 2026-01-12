package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/tool"
	"google.golang.org/genai"
	"nag763.com/recrudbagents/agents"
	"nag763.com/recrudbagents/utils"
)

func main() {
	ctx := context.Background()

	model, err := gemini.NewModel(ctx, utils.GetEnvOrDefault("MODEL", "gemini-2.5-flash-lite"), &genai.ClientConfig{
		Backend: genai.BackendVertexAI,
		Project: utils.GetEnvOrDefault("MODEL", "recrudb"),
	})

	mcpTools, err := utils.GetMcpTools(ctx, []string{"stores_offer_details", "delete_offer", "mark_communication_sent"})
	if err != nil {
		log.Fatalf("couldn't load mcp tools %v", err)
	}

	jobOfferMatcher, err := agents.NewJobOfferAgent(ctx, model)
	if err != nil {
		log.Fatalf("error while creating job offer agent %v", err)
	}

	reachOutAgent, err := agents.NewReachOutAgent(ctx, model)
	if err != nil {
		log.Fatalf("error while creating reach out agent %v", err)
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

		## reachout_agent

		Use it when requested to write a motivation letter, or to reach out the people mentionned in the offer.

		# Tools:
		## stores_offer_details

		Stores an offer details. This should be used only once an offer has already been analyzed by the joboffer_analyzer agent.

		You don't have to wait for user confirmation to store details, but in case some informations appear to be missing, indicate it to the user.

		## delete_offer

		Deletes an offer. Perform it using the UUID the job offer was created. Only do that with the job offers created in the session.

		## mark_communication_sent

		Marks a communication as sent, use it once a reach out mail has been written.
		`,
		Toolsets:  []tool.Toolset{mcpTools},
		SubAgents: []agent.Agent{jobOfferMatcher, reachOutAgent},
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
