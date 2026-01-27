package api

import "github.com/go-chi/cors"

var corsMiddleware = cors.Handler(cors.Options{
	AllowedOrigins:   []string{"https://*", "http://*"},
	AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	AllowCredentials: true,
	MaxAge:           300,
})
