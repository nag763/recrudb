package profile

import (
	"context"
	"log"
	"nag763/recrudb/api/internal/injection"
	"nag763/recrudb/api/internal/metadata"

	"github.com/mark3labs/mcp-go/mcp"
	"go.uber.org/zap"
)

func NewGetProfileTool(profileRepository Repository) injection.Tool {
	return injection.Tool{
		Tool: mcp.NewTool("get_profile",
			mcp.WithDescription("Get user profile"),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var metadata metadata.MetadataOnly
			if err := request.BindArguments(&metadata); err != nil {
				return mcp.NewToolResultError("ERROR WHILE ACCESSING DOC ID"), nil
			}

			doc, err := profileRepository.GetProfile(ctx, metadata.Metadata.UserId)
			if err != nil {
				return mcp.NewToolResultError("couldn't access profile"), nil
			}
			return mcp.NewToolResultStructuredOnly(doc), nil
		},
	}
}

func NewStoreProfileTool(profileRepository Repository) injection.Tool {
	return injection.Tool{
		Tool: mcp.NewTool("save_profile",
			mcp.WithDescription("Stores a CV analyzed details"),
			mcp.WithInputSchema[ProfileMcp](),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var doc ProfileMcpWithMeta
			if err := request.BindArguments(&doc); err != nil {
				return mcp.NewToolResultError("ERROR WHILE ACCESSING DOC ID"), nil
			}

			userId := doc.Metadata.UserId

			if _, err := profileRepository.StoreProfile(ctx, userId, doc.Document); err != nil {
				log.Printf("error while storing profile in firestore: %v\n", err)
				return mcp.NewToolResultError("ERROR WHILE STORING IN FIRESTORE"), nil
			}
			return mcp.NewToolResultStructured(doc, userId), nil
		},
	}
}

func NewGetOfferTool(profileRepository Repository, logger *zap.SugaredLogger) injection.Tool {
	return injection.Tool{
		Tool: mcp.NewTool("get_offer",
			mcp.WithDescription("Retrieves an offer details"),
			mcp.WithString("offerId"),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var metadata metadata.MetadataOnly
			if err := request.BindArguments(&metadata); err != nil {
				return mcp.NewToolResultError("ERROR WHILE ACCESSING DOC ID"), nil
			}

			offerId := request.GetString("offerId", "")
			if offerId == "" {
				return mcp.NewToolResultError("offer id is nil"), nil
			}

			doc, err := profileRepository.GetOffer(ctx, metadata.Metadata.UserId, offerId)
			if err != nil {
				return mcp.NewToolResultError("couldn't access profile"), nil
			}
			return mcp.NewToolResultStructuredOnly(doc), nil
		},
	}
}

func NewStoreOfferTool(profileRepository Repository, logger *zap.SugaredLogger) injection.Tool {
	return injection.Tool{
		Tool: mcp.NewTool("stores_offer",
			mcp.WithDescription("Stores offers details"),
			mcp.WithInputSchema[metadata.McpDocument[Offer]](),
		),
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var offer *metadata.WithMetadata[Offer]
			if err := request.BindArguments(&offer); err != nil {
				logger.Errorf("ERROR WHILE ACCESSING DOC ID")
				return mcp.NewToolResultError("ERROR WHILE ACCESSING DOC ID"), nil
			}

			id, err := profileRepository.StoreOffer(ctx, offer.Metadata.UserId, offer.Document)
			if err != nil {
				logger.Errorf("err ? %v", err)
				return mcp.NewToolResultError("ERROR WHILE STORING IN FIRESTORE"), nil
			}
			return mcp.NewToolResultText(id), nil
		},
	}
}
