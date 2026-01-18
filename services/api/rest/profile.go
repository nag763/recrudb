package rest

import (
	"encoding/json"
	"fmt"
	"nag763/recrudb/api/clients"
	"nag763/recrudb/api/domain"
	"nag763/recrudb/api/identity"
	"net/http"

	"github.com/go-chi/render"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sub, err := identity.GetSubFromRequest(ctx, r)
	if err != nil {
		render.Render(w, r, domain.ErrUnauthorizedResponse(fmt.Errorf("couldn't access sub")))
		return
	}

	client, err := clients.GetFirestoreClient()
	if err != nil {

		render.Render(w, r, domain.ErrInternalResponse(fmt.Errorf("couldn't access firestore")))
		return
	}

	dsnap, err := client.Collection("profile").Doc(sub).Get(ctx)
	if err != nil {
		render.Render(w, r, domain.ErrInvalidResponse(fmt.Errorf("profile %s not found", sub)))
		return
	}

	if err := json.NewEncoder(w).Encode(dsnap); err != nil {
		render.Render(w, r, domain.ErrInternalResponse(fmt.Errorf("couldn't encode json")))
		return
	}
}
