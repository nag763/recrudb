package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"main/middleware"
	"main/utils"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

var s *server.MCPServer
var client *firestore.Client
var err error

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
	CandidateScore        float32  `json:"candidate_score"`
	HowToApply            string   `json:"how_to_apply"`
}

type Communication struct {
	Initial  time.Time `firestore:"initial`
	Reminder time.Time `firstore:"reminder"`
	Last     time.Time `firestore:"last"`
}

func NewCommunication() Communication {
	initial := time.Now()
	reminder := initial.AddDate(0, 0, 7)
	last := initial.AddDate(0, 0, 21)
	return Communication{
		initial,
		reminder,
		last,
	}
}

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
	s.AddTool(
		mcp.NewTool("get_profile",
			mcp.WithDescription("Get user profile"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return mcp.NewToolResultText(ProfileText), nil
		},
	)

	s.AddTool(
		mcp.NewTool("stores_offer_details",
			mcp.WithDescription("Stores offers details"),
			mcp.WithInputSchema[OfferDetails](),
		),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			uuid := uuid.NewString()
			args := request.GetArguments()
			args["timestamp"] = time.Now()
			if _, err := client.Collection("offers").Doc(uuid).Set(ctx, request.GetArguments()); err != nil {
				log.Printf("error while storing job offer in firestore: %v\n", err)
				return mcp.NewToolResultError("ERROR WHILE STORING IN FIRESTORE"), nil
			}
			log.Printf("job offer stored under uuid %s\n", uuid)
			return mcp.NewToolResultText(uuid), nil
		},
	)

	s.AddTool(
		mcp.NewTool("delete_offer",
			mcp.WithDescription("Deletes an offer"),
			mcp.WithString("UUID", mcp.Required(), mcp.Description("The uuid of the offer to delete"))),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			uuid := request.GetString("UUID", "")
			if uuid == "" {
				return mcp.NewToolResultError("uuid param is mandatory"), nil
			}
			if _, err := client.Collection("offers").Doc(uuid).Delete(ctx); err != nil {
				log.Printf("error while deleting job offer in firestore: %v\n", err)
				return mcp.NewToolResultError(fmt.Sprintf("ERROR WHILE DELETING %s IN FIRESTORE", uuid)), nil
			}
			log.Printf("job %s offer deleted\n", uuid)
			return mcp.NewToolResultText(uuid), nil
		},
	)

	s.AddTool(
		mcp.NewTool("mark_communication_sent",
			mcp.WithDescription("Mark an offer with first communication sent"),
			mcp.WithString("UUID", mcp.Required(), mcp.Description("The uuid of the offer to mark"))),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			uuid := request.GetString("UUID", "")
			if uuid == "" {
				return mcp.NewToolResultError("uuid param is mandatory"), nil
			}

			communication := NewCommunication()
			if _, err := client.Collection("offers").Doc(uuid).Set(ctx, communication, firestore.MergeAll); err != nil {
				log.Printf("error while deleting job offer in firestore: %v\n", err)
				return mcp.NewToolResultError(fmt.Sprintf("ERROR WHILE DELETING %s IN FIRESTORE", uuid)), nil
			}
			log.Printf("job %s offer deleted\n", uuid)
			return mcp.NewToolResultText(uuid), nil
		},
	)

	s.AddTool(
		mcp.NewTool("store_cv",
			mcp.WithDescription("Stores a CV analyzed details"),
			mcp.WithInputSchema[struct {
				Firstname   string
				Lastname    string
				Contact     Contact `json:",omitempty"`
				Skills      []string
				LookingFor  string
				Experiences []struct {
					StartDate time.Time
					EndDate   time.Time
					Company   string
					Misison   string
					Role      string
					Skills    []string
				}
				Languages []string
				Socials   []struct {
					Name string
					Link string
				}
				Location Address `json:",omitempty"`
			}](),
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
		},
	)

}

func main() {
	defer client.Close()

	httpServer := server.NewSSEServer(s)

	if err := httpServer.Start(utils.GetEnvOrDefault("ADRESS", ":3000")); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
