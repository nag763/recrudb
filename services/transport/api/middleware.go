package api

import (
	"context"
	"nag763/recrudb/api/internal/pkg/identity"
	"nag763/recrudb/api/internal/response"
	"net/http"

	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

var corsMiddleware = cors.Handler(cors.Options{
	AllowedOrigins:   []string{"https://*", "http://*"},
	AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	AllowCredentials: true,
	MaxAge:           300,
})

func authenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sub, err := identity.GetSubFromRequest(ctx, r.Header)
		if err != nil {
			render.Render(w, r, response.Unauthorized(err))
			return
		}

		ctx = context.WithValue(r.Context(), "sub", sub)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
