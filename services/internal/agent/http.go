package agent

import (
	"fmt"
	"nag763/recrudb/api/internal/agent/firestore"
	"nag763/recrudb/api/internal/pkg/identity"
	"nag763/recrudb/api/internal/request"
	"nag763/recrudb/api/internal/response"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"google.golang.org/adk/session"
)

func interactionHttpHandler(runner *Runner) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sub, err := identity.GetSubFromRequest(r.Context(), r.Header)
		if err != nil {
			render.Render(w, r, response.Unauthorized(fmt.Errorf("no sub found : %w", err)))
			return
		}

		httpRequest := &request.HttpRequest{}
		switch header := r.Header.Get("Content-Type"); {
		case strings.Contains(header, "multipart/form-data"):
			httpRequest, err = request.FromMultipart(w, r)
		default:
			err = render.Bind(r, httpRequest)
		}
		if err != nil {
			render.Render(w, r, response.InvalidRequest(err))
			return
		}

		resp, err := runner.Run(RunParams{
			Input:        httpRequest.ToGenAiContent(),
			UserId:       sub,
			ClearSession: httpRequest.ClearSession,
			SessionId:    httpRequest.SessionId,
		})

		if err != nil {
			render.Render(w, r, response.InternalResponse(fmt.Errorf("issue with runner %v", err)))
			return
		}

		render.Render(w, r, response.Success(resp))
	}
}

func sessionListHttpHandler(runner *Runner, sessionService *firestore.FirestoreSessionService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sub, err := identity.GetSubFromRequest(r.Context(), r.Header)
		if err != nil {
			render.Render(w, r, response.Unauthorized(fmt.Errorf("no sub found : %w", err)))
			return
		}

		sessions, err := sessionService.ListSessionShort(r.Context(), &session.ListRequest{
			AppName: runner.appName,
			UserID:  sub,
		})
		if err != nil {
			render.Render(w, r, response.InternalResponse(err))
		}
		render.Render(w, r, response.Success(sessions))
	}
}
