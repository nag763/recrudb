package api

import (
	"context"
	"nag763/recrudb/api/internal/profile"
	"nag763/recrudb/api/internal/response"
	"nag763/recrudb/api/internal/subrouter"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/fx"
)

type RouterParams struct {
	fx.In
	SubRouters []subrouter.SubRouter `group:"subrouters"`
}

func NewHttpServer(lc fx.Lifecycle, p RouterParams, profileHandler *profile.ProfileHandler) *chi.Mux {

	r := chi.NewRouter()

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, response.NotFoundResponse())
	})

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger, middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		for _, sub := range p.SubRouters {
			r.Mount(sub.Path, sub.Router)
		}
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go http.ListenAndServe(":8000", r)
			return nil
		},
	})
	return r

}
