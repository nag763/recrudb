package agent

import (
	"fmt"
	"nag763/recrudb/api/internal/pkg/identity"
	"nag763/recrudb/api/internal/request"
	"nag763/recrudb/api/internal/response"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type AgentRouter struct {
	Router chi.Router
}

func NewAgentRouter(agentService *AgentService) *AgentRouter {
	r := chi.NewMux()

	for path, runner := range agentService.Runners {
		r.Post("/"+path, func(w http.ResponseWriter, r *http.Request) {
			sub, err := identity.GetSubFromRequest(r.Context(), r.Header)
			if err != nil {
				render.Render(w, r, response.Unauthorized(fmt.Errorf("no sub found")))
				return
			}

			httpRequest := &request.HttpRequest{}
			if err := render.Bind(r, httpRequest); err != nil {
				render.Render(w, r, response.CleintInvalidRequest(err))
				return
			}

			resp, err := runner.Run(r.Context(), RunParams{
				Input:        httpRequest.Content,
				UserId:       sub,
				ClearSession: httpRequest.ClearSession,
			})
			if err != nil {
				render.Render(w, r, response.InternalResponse(fmt.Errorf("issue with runner %v", err)))
				return
			}

			render.Render(w, r, response.Success(resp))
		})
	}

	return &AgentRouter{
		Router: r,
	}
}
