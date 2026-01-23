package main

import (
	"nag763/recrudb/api/internal/db"
	"nag763/recrudb/api/internal/offer"
	"nag763/recrudb/api/internal/pkg/logger"
	"nag763/recrudb/api/internal/profile"
	"nag763/recrudb/api/transport/mcp"
	"nag763/recrudb/api/transport/mcp/middleware"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		logger.Module,
		db.Module,
		profile.Module,
		offer.Module,
		middleware.Module,
		mcp.Module,
	).Run()
}
