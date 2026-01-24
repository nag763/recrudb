package offer

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func NewStoreOfferTool(offerRepository *OfferRepository) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("stores_offer",
			mcp.WithDescription("Stores offers details"),
			mcp.WithInputSchema[Offer](),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var offer *Offer
			if err := request.BindArguments(&offer); err != nil {
				return mcp.NewToolResultError("ERROR WHILE ACCESSING DOC ID"), nil
			}

			if _, err := offerRepository.StoreOffer(ctx, *offer); err != nil {
				return mcp.NewToolResultError("ERROR WHILE STORING IN FIRESTORE"), nil
			}
			return mcp.NewToolResultText("OK"), nil
		}
}
