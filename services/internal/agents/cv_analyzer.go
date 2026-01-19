package agents

import (
	"context"
	_ "embed"
	"nag763/recrudb/api/internal/clients"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/tool"
)

//go:embed prompts/cv_analyzer.txt
var cvAnalyzer string

func NewCvAnalyzer(ctx context.Context, model model.LLM) (agent.Agent, error) {

	tools, err := clients.GetMcpTools(ctx, []string{"store_cv"})
	if err != nil {
		return nil, err
	}

	agent, err := llmagent.New(llmagent.Config{
		Name:        "cv_analyzer",
		Model:       model,
		Description: "Analyze CV and store details.",
		Instruction: cvAnalyzer,
		Toolsets:    []tool.Toolset{tools},
	})
	if err != nil {
		return nil, err
	}

	return agent, nil
}
