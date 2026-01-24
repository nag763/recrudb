package subrouter

import (
	"github.com/go-chi/chi/v5"
)

type SubRouter struct {
	Path   string
	Router chi.Router
}
