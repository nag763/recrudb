package main

import (
	"context"
	"fmt"
	"log"
	"nag763/recrudb/api/internal/utils"
	"nag763/recrudb/api/transport/mcp/middleware"
	"nag763/recrudb/api/transport/mcp/tools"

	"cloud.google.com/go/firestore"
	"github.com/mark3labs/mcp-go/server"
)

var s *server.MCPServer
var client *firestore.Client
var err error

func init() {
	ctx := context.Background()

	client, err = firestore.NewClient(ctx, utils.GetEnvOrDefault("FIRESTORE_DB", "recrudb"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Create a new MCP server
	s = server.NewMCPServer(
		"RecruDB Agent Tools",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithToolHandlerMiddleware(middleware.AuthorizationHook),
	)

	// Add tool handler
	s.AddTool(tools.NewStoreOfferTool(client))
	s.AddTool(tools.NewGetProfileTool(ctx, client))
	s.AddTool(tools.NewDeleteOfferTool(client))
	s.AddTool(tools.NewMarkCommunicationSentTool(client))
	s.AddTool(tools.NewStoreProfileTool(client))

}

func main() {
	defer client.Close()

	httpServer := server.NewSSEServer(s)

	if err := httpServer.Start(utils.GetEnvOrDefault("ADRESS", ":3000")); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
