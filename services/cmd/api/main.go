package main

import (
	"nag763/recrudb/api/internal/agent"
	"nag763/recrudb/api/internal/db"
	"nag763/recrudb/api/internal/profile"
	"nag763/recrudb/api/transport/api"

	"go.uber.org/fx"
)

func main() {

	fx.New(
		db.Module,
		profile.Module,
		agent.Module,
		api.Module,
	).Run()
}
