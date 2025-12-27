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

type Contact struct {
	Email       string `json:"email,omitempty"`
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"number,omitempty"`
}

type Address struct {
	AddressLine1 string `json:"line1,omitempty"`
	AddressLine2 string `json:"line2,omitempty"`
	AddressLine3 string `json:"line3,omitempty"`
	City         string `json:"city,omitempty"`
	Country      string `json:"country,omitempty"`
}

type OfferDetails struct {
	EnterpriseName        string   `json:"enterpriseName"`
	Contact               Contact  `json:"contact"`
	Address               Address  `json:"address"`
	TechnicalRequirements []string `json:"tech_req"`
}

func init() {
	viper.SetDefault("ADRESS", ":3000")
	viper.AutomaticEnv()

	// Create a new MCP server
	s = server.NewMCPServer(
		"RecruDB Agent Tools",
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

	s.AddTool(
		mcp.NewTool("stores_offer_details",
			mcp.WithDescription("Stores offers details"),
			mcp.WithInputSchema[OfferDetails](),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			log.Printf("tool called w/ %v", request)
			return mcp.NewToolResultText("OK"), nil
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
