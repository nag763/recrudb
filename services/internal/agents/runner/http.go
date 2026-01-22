package runner

import (
	"fmt"
	"nag763/recrudb/api/internal/pkg/identity"
	"nag763/recrudb/api/internal/request"
	"nag763/recrudb/api/internal/response"
	"net/http"
	"sync"

	"github.com/go-chi/render"
)

func SessionlessHttpRunner(initalizer AgentInitializer) func(w http.ResponseWriter, r *http.Request) {

	runnerInitalizer := sync.OnceValues(func() (*AgentRunner, error) {
		return CreateAgentRunner(initalizer)
	})

	return func(w http.ResponseWriter, r *http.Request) {

		sub, err := identity.GetSubFromRequest(r.Context(), r.Header)
		if err != nil {
			render.Render(w, r, response.ErrUnauthorizedResponse(fmt.Errorf("no sub found")))
			return
		}

		httpRequest := &request.HttpRequest{}
		if err := render.Bind(r, httpRequest); err != nil {
			render.Render(w, r, response.ErrInvalidResponse(err))
			return
		}

		runner, err := runnerInitalizer()
		if err != nil {
			render.Render(w, r, response.ErrInternalResponse(fmt.Errorf("couldn't start runner %v", err)))
			return
		}

		resp, err := runner.RunSessionless(httpRequest.Content, sub)
		if err != nil {
			render.Render(w, r, response.ErrInternalResponse(fmt.Errorf("issue with runner %v", err)))
			return
		}

		render.Render(w, r, response.SuccessResponse(resp))
	}
}
