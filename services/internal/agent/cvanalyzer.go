package agent

import (
	_ "embed"
	"nag763/recrudb/api/internal/injection"

	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/tool"
	"google.golang.org/genai"
)

//go:embed prompts/cv_analyzer.txt
var cvAnalyzer string

func NewCvAnalyzer(agentConfig *AgentConfig) (*injection.HttpAgent, error) {

	allowedTools := []string{"store_cv"}
	tools, err := getMcpTools(allowedTools)
	if err != nil {
		return nil, err
	}

	agent, err := llmagent.New(llmagent.Config{
		Name:                "cv_analyzer",
		Model:               agentConfig.Model,
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

	return &injection.HttpAgent{
		Path:  "/cvAnalyzer",
		Agent: agent,
	}, nil
}
