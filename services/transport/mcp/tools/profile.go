package tools

import (
	"context"
	"log"
	"nag763/recrudb/api/internal/domain"
	"nag763/recrudb/api/internal/identity"

	"cloud.google.com/go/firestore"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func NewGetProfileTool(ctx context.Context, client *firestore.Client) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_profile",
			mcp.WithDescription("Get user profile"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			sub, err := identity.GetSubFromRequest(ctx, request.Header)
			if err != nil {
				return mcp.NewToolResultError("couldn't extract sub"), nil
			}

			doc, err := domain.GetProfile(ctx, client, sub)
			if err != nil {
				return mcp.NewToolResultError("couldn't access profile"), nil
			}
			return mcp.NewToolResultStructuredOnly(doc), nil
		}
}

func NewStoreProfileTool(client *firestore.Client) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("store_cv",
			mcp.WithDescription("Stores a CV analyzed details"),
			mcp.WithInputSchema[domain.Profile](),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var doc domain.Profile
			if err := request.BindArguments(&doc); err != nil {
				return mcp.NewToolResultError("ERROR WHILE ACCESSING DOC ID"), nil
			}

			sub, err := identity.GetSubFromRequest(ctx, request.Header)
			if err != nil {
				return mcp.NewToolResultError("COULDN'T ACCESS CURRENT USER"), err
			}

			if _, err := client.Collection("profile").Doc(sub).Set(ctx, doc); err != nil {
				log.Printf("error while storing profile in firestore: %v\n", err)
				return mcp.NewToolResultError("ERROR WHILE STORING IN FIRESTORE"), nil
			}
			return mcp.NewToolResultStructured(map[string]any{"id": sub, "doc": doc}, sub), nil
		}
}
