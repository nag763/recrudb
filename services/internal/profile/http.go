package profile

import (
	"nag763/recrudb/api/internal/pkg/identity"
	"nag763/recrudb/api/internal/response"
	"net/http"

	"github.com/go-chi/render"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProfileHandler struct {
	profileRepository *ProfileRepository
}

func NewProfileHandler(profileRepository *ProfileRepository) *ProfileHandler {
	return &ProfileHandler{
		profileRepository,
	}
}

func (h *ProfileHandler) HandleGetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sub, err := identity.GetSubFromRequest(ctx, r.Header)
	if err != nil {
		render.Render(w, r, response.Unauthorized(err))
		return
	}

	doc, err := h.profileRepository.GetProfile(ctx, sub)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			render.Render(w, r, response.NoContent())
			return
		}
		render.Render(w, r, response.InternalResponse(err))
		return
	}

	render.Render(w, r, response.Success(doc))
}

func (h *ProfileHandler) HandleGetProfileOffers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sub, err := identity.GetSubFromRequest(ctx, r.Header)
	if err != nil {
		render.Render(w, r, response.Unauthorized(err))
		return
	}

	doc, err := h.profileRepository.GetOffers(ctx, sub)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			render.Render(w, r, response.NoContent())
			return
		}
		render.Render(w, r, response.InternalResponse(err))
		return
	}

	render.Render(w, r, response.Success(doc))
}
