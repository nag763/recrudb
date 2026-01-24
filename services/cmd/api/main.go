package main

import (
	"nag763/recrudb/api/internal/agent"
	"nag763/recrudb/api/internal/db"
	"nag763/recrudb/api/internal/profile"
	"nag763/recrudb/api/transport/api"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		db.Module,
		profile.Module,
		agent.Module,
		api.Module,
		fx.Invoke(func(*chi.Mux) {}),
	).Run()
}
