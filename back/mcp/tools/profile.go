package tools

import (
	"context"
	"log"
	"main/domain"
	"main/utils"

	"cloud.google.com/go/firestore"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

//go:embed profile.txt
var ProfileText string

func NewGetProfileTool(ctx context.Context, client *firestore.Client) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_profile",
			mcp.WithDescription("Get user profile"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

			bearer, err := utils.ExtractBearer(request)
			if err != nil {
				return mcp.NewToolResultError("couldn't extract authorization"), nil
			}

			sub, err := utils.GetSub(ctx, bearer)
			if err != nil {
				return mcp.NewToolResultError("couldn't extract sub"), nil
			}

			doc, err := client.Collection("profile").Doc(sub).Get(ctx)
			if err != nil {
				return mcp.NewToolResultError("couldn't access profile"), nil
			}
			return mcp.NewToolResultStructuredOnly(doc.Data()), nil
		}
}

func NewStoreProfileTool(client *firestore.Client) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("store_cv",
			mcp.WithDescription("Stores a CV analyzed details"),
			mcp.WithInputSchema[domain.Profile](),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

			bearer, err := utils.ExtractBearer(request)
			if err != nil {
				return mcp.NewToolResultError("couldn't extract authorization"), nil
			}

			sub, err := utils.GetSub(ctx, bearer)
			if err != nil {
				return mcp.NewToolResultError("couldn't extract sub"), nil
			}

			if _, err := client.Collection("profile").Doc(sub).Set(ctx, request.GetArguments()); err != nil {
				log.Printf("error while storing profile in firestore: %v\n", err)
				return mcp.NewToolResultError("ERROR WHILE STORING IN FIRESTORE"), nil
			}
			return mcp.NewToolResultText(sub), nil
		}
}
