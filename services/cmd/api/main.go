package main

import (
	"log"
	"nag763/recrudb/api/internal/agents"
	"nag763/recrudb/api/internal/agents/runner"
	"nag763/recrudb/api/internal/clients"
	"nag763/recrudb/api/internal/domain"
	"nag763/recrudb/api/transport/api"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type HandlerFuncWithClient[T any] func(w http.ResponseWriter, r *http.Request, client *T)

func injectClient[T firestore.Client](fn HandlerFuncWithClient[T], client *T) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, client)
	}
}

func main() {

	firestoreClient, err := clients.GetFirestoreClient()
	if err != nil {
		log.Fatalf("couldn't acces firestore client %v", err)
	}
	r := chi.NewRouter()

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, domain.ErrNotFoundResponse())
	})
	r.Use(middleware.Logger, middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/agents", func(r chi.Router) {
			r.Post("/jobOffer", runner.SessionlessHttpRunner(agents.NewJobOfferAgent))
			r.Post("/reachOut", runner.SessionlessHttpRunner(agents.NewReachOutAgent))
			r.Post("/cvAnalyzer", runner.SessionlessHttpRunner(agents.NewCvAnalyzer))
		})
		r.Get("/profile/me", injectClient(api.GetProfile, firestoreClient))
	})

	if err := http.ListenAndServe(":8000", r); err != nil {
		panic(err)
	}

}
