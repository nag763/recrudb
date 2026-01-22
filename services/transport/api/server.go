package api

import (
	"context"
	"nag763/recrudb/api/internal/agents"
	"nag763/recrudb/api/internal/agents/runner"
	"nag763/recrudb/api/internal/domain"
	"nag763/recrudb/api/internal/profile"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/fx"
)

func NewHttpServer(lc fx.Lifecycle, profileHandler *profile.ProfileHandler) *chi.Mux {

	r := chi.NewRouter()

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, domain.ErrNotFoundResponse())
	})

	r.Use(middleware.Timeout(time.Second * 30))
	r.Use(middleware.Logger, middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/agents", func(r chi.Router) {
			r.Post("/jobOffer", runner.SessionlessHttpRunner(agents.NewJobOfferAgent))
			r.Post("/reachOut", runner.SessionlessHttpRunner(agents.NewReachOutAgent))
			r.Post("/cvAnalyzer", runner.SessionlessHttpRunner(agents.NewCvAnalyzer))
		})
		r.Get("/profile/me", profileHandler.HandleGetProfile)
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go http.ListenAndServe(":8000", r)
			return nil
		},
	})
	return r

}
