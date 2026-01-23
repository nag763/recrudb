package agent

import (
	"context"
	_ "embed"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model"
	"google.golang.org/adk/tool"
	"google.golang.org/genai"
)

//go:embed prompts/cv_analyzer.txt
var cvAnalyzer string

func NewCvAnalyzer(ctx context.Context, model model.LLM) (agent.Agent, error) {

	allowedTools := []string{"store_cv"}
	tools, err := getMcpTools(ctx, allowedTools)
	if err != nil {
		return nil, err
	}

	agent, err := llmagent.New(llmagent.Config{
		Name:                "cv_analyzer",
		Model:               model,
		Description:         "Analyze CV and store details.",
		Instruction:         cvAnalyzer,
		Toolsets:            []tool.Toolset{tools},
		BeforeToolCallbacks: []llmagent.BeforeToolCallback{injectMetadata},
		GenerateContentConfig: &genai.GenerateContentConfig{
			ToolConfig: &genai.ToolConfig{
				FunctionCallingConfig: restrictConfigTo(allowedTools),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return agent, nil
}
