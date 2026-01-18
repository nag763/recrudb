package runner

import (
	"fmt"
	"nag763/recrudb/api/internal/domain"
	"net/http"
	"sync"

	"github.com/go-chi/render"
)

func SessionlessHttpRunner(initalizer AgentInitializer) func(w http.ResponseWriter, r *http.Request) {

	runnerInitalizer := sync.OnceValues(func() (*AgentRunner, error) {
		return CreateAgentRunner(initalizer)
	})

	return func(w http.ResponseWriter, r *http.Request) {

		httpRequest := &domain.HttpRequest{}
		if err := render.Bind(r, httpRequest); err != nil {
			render.Render(w, r, domain.ErrInvalidResponse(err))
			return
		}

		runner, err := runnerInitalizer()
		if err != nil {
			render.Render(w, r, domain.ErrInternalResponse(fmt.Errorf("couldn't start runner")))
			return
		}

		resp, err := runner.RunSessionless(httpRequest.Content)
		if err != nil {
			render.Render(w, r, domain.ErrInternalResponse(fmt.Errorf("issue with runner %v", runner)))
			return
		}

		render.Render(w, r, domain.SuccessResponse(resp))
	}
}
