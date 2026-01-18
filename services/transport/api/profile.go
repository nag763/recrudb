package api

import (
	"nag763/recrudb/api/internal/domain"
	"nag763/recrudb/api/internal/identity"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/go-chi/render"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetProfile(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
	sub, err := identity.GetSubFromRequest(r.Context(), r.Header)
	if err != nil {
		render.Render(w, r, domain.ErrUnauthorizedResponse(err))
		return
	}

	doc, err := domain.GetProfile(r.Context(), client, sub)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			render.Render(w, r, domain.NoContentResponse())
			return
		}
		render.Render(w, r, domain.ErrInternalResponse(err))
		return
	}

	render.Render(w, r, domain.SuccessResponse(doc))
}
