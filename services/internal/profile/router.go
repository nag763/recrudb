package profile

import (
	"nag763/recrudb/api/internal/subrouter"

	"github.com/go-chi/chi/v5"
)

func NewProfileRouter(profileHandler *ProfileHandler) subrouter.SubRouter {
	r := chi.NewMux()
	r.Get("/me", profileHandler.HandleGetProfile)
	return subrouter.SubRouter{
		Router: r,
		Path:   "/profile",
	}
}
