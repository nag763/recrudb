package profile

import (
	"nag763/recrudb/api/internal/injection"

	"github.com/go-chi/chi/v5"
)

func NewProfileRouter(profileHandler *ProfileHandler) injection.SubRouter {
	r := chi.NewMux()

	r.Get("/me", profileHandler.HandleGetProfile)
	r.Get("/me/offers", profileHandler.HandleGetProfileOffers)

	return injection.SubRouter{
		Router: r,
		Path:   "/profile",
	}
}
