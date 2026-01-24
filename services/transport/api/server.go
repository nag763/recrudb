package api

import (
	"context"
	"nag763/recrudb/api/internal/injection"
	"nag763/recrudb/api/internal/response"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/fx"
)

type RouterParams struct {
	fx.In
	Middlewares []func(http.Handler) http.Handler `group:"api_middlewares"`
	SubRouters  []injection.SubRouter             `group:"subrouters"`
}

func NewHttpServer(lc fx.Lifecycle, p RouterParams) *chi.Mux {

	r := chi.NewRouter()

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, response.NotFoundResponse())
	})

	r.Use(p.Middlewares...)

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
