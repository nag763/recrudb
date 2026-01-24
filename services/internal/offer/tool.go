package offer

import (
	"context"
	"nag763/recrudb/api/internal/injection"

	"github.com/mark3labs/mcp-go/mcp"
)

func NewStoreOfferTool(offerRepository *OfferRepository) injection.Tool {
	return injection.Tool{
		Tool: mcp.NewTool("stores_offer",
			mcp.WithDescription("Stores offers details"),
			mcp.WithInputSchema[Offer](),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var offer *Offer
			if err := request.BindArguments(&offer); err != nil {
				return mcp.NewToolResultError("ERROR WHILE ACCESSING DOC ID"), nil
			}

			if _, err := offerRepository.StoreOffer(ctx, *offer); err != nil {
				return mcp.NewToolResultError("ERROR WHILE STORING IN FIRESTORE"), nil
			}
			return mcp.NewToolResultText("OK"), nil
		},
	}
}
