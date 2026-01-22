package profile

import (
	"context"
	"log"
	"nag763/recrudb/api/internal/domain"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type ProfileTool struct {
	profileRepository *ProfileRepository
}

func NewProfileTool(profileRepository *ProfileRepository) *ProfileTool {
	return &ProfileTool{
		profileRepository,
	}
}

func (t *ProfileTool) NewGetProfileTool() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_profile",
			mcp.WithDescription("Get user profile"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var metadata domain.MetadataOnly
			if err := request.BindArguments(&metadata); err != nil {
				return mcp.NewToolResultError("ERROR WHILE ACCESSING DOC ID"), nil
			}

			doc, err := t.profileRepository.GetProfile(ctx, metadata.Metadata.UserId)
			if err != nil {
				return mcp.NewToolResultError("couldn't access profile"), nil
			}
			return mcp.NewToolResultStructuredOnly(doc), nil
		}
}

func (t *ProfileTool) NewStoreProfileTool() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("store_cv",
			mcp.WithDescription("Stores a CV analyzed details"),
			mcp.WithInputSchema[ProfileMcp](),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var doc ProfileMcpWithMeta
			if err := request.BindArguments(&doc); err != nil {
				return mcp.NewToolResultError("ERROR WHILE ACCESSING DOC ID"), nil
			}

			userId := doc.Metadata.UserId

			if _, err := t.profileRepository.StoreProfile(ctx, userId, doc.Document); err != nil {
				log.Printf("error while storing profile in firestore: %v\n", err)
				return mcp.NewToolResultError("ERROR WHILE STORING IN FIRESTORE"), nil
			}
			return mcp.NewToolResultStructured(doc, userId), nil
		}
}
