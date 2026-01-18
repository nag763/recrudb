package main

import (
	"nag763/recrudb/api/agents"
	"nag763/recrudb/api/agents/runner"
	"nag763/recrudb/api/domain"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, domain.ErrNotFoundResponse())
	})
	r.Use(middleware.Logger, middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/agents", func(r chi.Router) {
			r.Post("/jobOffer", runner.SessionlessHttpRunner(agents.NewJobOfferAgent))
			r.Post("/reachOut", runner.SessionlessHttpRunner(agents.NewReachOutAgent))
			r.Post("/cvAnalyer", runner.SessionlessHttpRunner(agents.NewCvAnalyzer))
		})
	})

	if err := http.ListenAndServe(":8000", r); err != nil {
		panic(err)
	}

}
