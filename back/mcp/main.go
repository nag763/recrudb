package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/viper"
)

var s *server.MCPServer

//go:embed profile.txt
var ProfileText string

func init() {
	viper.SetDefault("ADRESS", ":3000")
	viper.AutomaticEnv()

	// Create a new MCP server
	s = server.NewMCPServer(
		"MCP",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	// Add tool handler
	s.AddTool(
		mcp.NewTool("get_profile",
			mcp.WithDescription("Get user profile"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			log.Printf("Tool called")
			return mcp.NewToolResultText(ProfileText), nil
		},
	)
}

func main() {
	httpServer := server.NewStreamableHTTPServer(s)
	// Start the stdio server
	if err := httpServer.Start(viper.GetString("ADRESS")); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
