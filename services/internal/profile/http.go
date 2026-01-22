package profile

import (
	"nag763/recrudb/api/internal/domain"
	"nag763/recrudb/api/internal/identity"
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
		render.Render(w, r, domain.ErrUnauthorizedResponse(err))
		return
	}

	doc, err := h.profileRepository.GetProfile(ctx, sub)
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
