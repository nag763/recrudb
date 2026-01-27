package main

import (
	"nag763/recrudb/api/transport/api"
	"nag763/recrudb/api/transport/mcp"

	"github.com/go-chi/chi/v5"
	"github.com/mark3labs/mcp-go/server"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		mcp.Module,
		api.Module,
		fx.Invoke(func(*server.StreamableHTTPServer) {}),
		fx.Invoke(func(*chi.Mux) {}),
	).Run()
}
