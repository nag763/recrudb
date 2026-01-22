package offer

import (
	"context"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type OfferTool struct {
	offerRepository *OfferRepository
}

func NewOfferTool(offerRepository *OfferRepository) *OfferTool {
	return &OfferTool{
		offerRepository,
	}
}

func (t *OfferTool) NewStoreOfferTool() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("stores_offer",
			mcp.WithDescription("Stores offers details"),
			mcp.WithInputSchema[Offer](),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var offer *Offer
			if err := request.BindArguments(&offer); err != nil {
				return mcp.NewToolResultError("ERROR WHILE ACCESSING DOC ID"), nil
			}

			if _, err := t.offerRepository.StoreOffer(ctx, *offer); err != nil {
				log.Printf("error while storing job offer in firestore: %v\n", err)
				return mcp.NewToolResultError("ERROR WHILE STORING IN FIRESTORE"), nil
			}
			return mcp.NewToolResultText("OK"), nil
		}
}
