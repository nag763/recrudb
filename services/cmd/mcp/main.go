package main

import (
	"nag763/recrudb/api/internal/db"
	"nag763/recrudb/api/internal/offer"
	"nag763/recrudb/api/internal/profile"
	"nag763/recrudb/api/transport/mcp"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		db.Module,
		profile.Module,
		offer.Module,
		mcp.Module,
	).Run()
}
