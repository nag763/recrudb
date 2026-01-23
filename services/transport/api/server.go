package api

import (
	"context"
	"nag763/recrudb/api/internal/agent"
	"nag763/recrudb/api/internal/profile"
	"nag763/recrudb/api/internal/response"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/fx"
)

func NewHttpServer(lc fx.Lifecycle, profileHandler *profile.ProfileHandler, agentRouter *agent.AgentRouter) *chi.Mux {

	r := chi.NewRouter()

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, response.NotFoundResponse())
	})

	// r.Use(middleware.Timeout(time.Minute))

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger, middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/agents", agentRouter.Router)
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
