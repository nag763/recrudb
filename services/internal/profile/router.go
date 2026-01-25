package profile

import (
	"nag763/recrudb/api/internal/injection"

	"github.com/go-chi/chi/v5"
)

func NewProfileRouter(profileHandler *ProfileHandler) injection.SubRouter {
	r := chi.NewMux()

	r.Get("/me", profileHandler.HandleGetProfile)
	r.Delete("/me", profileHandler.HandleDeleteProfile)
	r.Get("/me/offers", profileHandler.HandleGetProfileOffers)
	r.Get("/me/offers/{id}", profileHandler.HandleGetOffer)
	r.Delete("/me/", profileHandler.HandleDeleteProfile)
	r.Delete("/me/offers/{id}", profileHandler.HandleDeleteOffer)

	return injection.SubRouter{
		Router: r,
		Path:   "/profile",
	}
}
