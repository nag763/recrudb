package agent

import (
	"fmt"
	"nag763/recrudb/api/internal/pkg/identity"
	"nag763/recrudb/api/internal/request"
	"nag763/recrudb/api/internal/response"
	"net/http"

	"github.com/go-chi/render"
)

func httpHandler(runner *Runner) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sub, err := identity.GetSubFromRequest(r.Context(), r.Header)
		if err != nil {
			render.Render(w, r, response.Unauthorized(fmt.Errorf("no sub found : %w", err)))
			return
		}

		httpRequest := &request.HttpRequest{}
		if err := render.Bind(r, httpRequest); err != nil {
			render.Render(w, r, response.InvalidRequest(err))
			return
		}

		resp, err := runner.Run(RunParams{
			Input:        httpRequest.Content,
			UserId:       sub,
			ClearSession: httpRequest.ClearSession,
		})
		if err != nil {
			render.Render(w, r, response.InternalResponse(fmt.Errorf("issue with runner %v", err)))
			return
		}

		render.Render(w, r, response.Success(resp))
	}
}
