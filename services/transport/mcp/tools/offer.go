package tools

import (
	"context"
	"fmt"
	"log"
	"nag763/recrudb/api/internal/domain"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func NewStoreOfferTool(client *firestore.Client) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("stores_offer",
			mcp.WithDescription("Stores offers details"),
			mcp.WithInputSchema[domain.Offer](),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			uuid := uuid.NewString()

			var offer *domain.Offer
			if err := request.BindArguments(&offer); err != nil {
				return mcp.NewToolResultError("ERROR WHILE ACCESSING DOC ID"), nil
			}

			if _, err := client.Collection("offers").Doc(uuid).Set(ctx, offer); err != nil {
				log.Printf("error while storing job offer in firestore: %v\n", err)
				return mcp.NewToolResultError("ERROR WHILE STORING IN FIRESTORE"), nil
			}
			return mcp.NewToolResultText(uuid), nil
		}
}

func NewDeleteOfferTool(client *firestore.Client) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("mark_communication_sent",
			mcp.WithDescription("Mark an offer with first communication sent"),
			mcp.WithString("UUID", mcp.Required(), mcp.Description("The uuid of the offer to mark"))),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			uuid := request.GetString("UUID", "")
			if uuid == "" {
				return mcp.NewToolResultError("uuid param is mandatory"), nil
			}

			communication := domain.NewCommunication()
			if _, err := client.Collection("offers").Doc(uuid).Set(ctx, communication, firestore.MergeAll); err != nil {
				log.Printf("error while deleting job offer in firestore: %v\n", err)
				return mcp.NewToolResultError(fmt.Sprintf("ERROR WHILE DELETING %s IN FIRESTORE", uuid)), nil
			}
			log.Printf("job %s offer deleted\n", uuid)
			return mcp.NewToolResultText(uuid), nil
		}
}

func NewMarkCommunicationSentTool(client *firestore.Client) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("mark_communication_sent",
			mcp.WithDescription("Mark an offer with first communication sent"),
			mcp.WithString("UUID", mcp.Required(), mcp.Description("The uuid of the offer to mark"))),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			uuid := request.GetString("UUID", "")
			if uuid == "" {
				return mcp.NewToolResultError("uuid param is mandatory"), nil
			}

			communication := domain.NewCommunication()
			if _, err := client.Collection("offers").Doc(uuid).Set(ctx, communication, firestore.MergeAll); err != nil {
				log.Printf("error while deleting job offer in firestore: %v\n", err)
				return mcp.NewToolResultError(fmt.Sprintf("ERROR WHILE DELETING %s IN FIRESTORE", uuid)), nil
			}
			log.Printf("job %s offer deleted\n", uuid)
			return mcp.NewToolResultText(uuid), nil
		}
}
