package config

import "google.golang.org/genai"

func RestrictConfigTo(tools []string) *genai.FunctionCallingConfig {
	return &genai.FunctionCallingConfig{
		AllowedFunctionNames: tools,
		Mode:                 genai.FunctionCallingConfigModeValidated,
	}
}
