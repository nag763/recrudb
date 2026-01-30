package agent

import (
	"fmt"
	"io"
	"nag763/recrudb/api/internal/pkg/identity"
	"nag763/recrudb/api/internal/request"
	"nag763/recrudb/api/internal/response"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"google.golang.org/genai"
)

func parseRequest(w http.ResponseWriter, r *http.Request) (*genai.Content, error) {
	switch header := r.Header.Get("Content-Type"); {

	case header == "application/json":
		httpRequest := &request.HttpRequest{}
		if err := render.Bind(r, httpRequest); err != nil {
			render.Render(w, r, response.InvalidRequest(err))
			return nil, err
		}
		return genai.NewContentFromText(httpRequest.Data, genai.RoleUser), nil

	case strings.Contains(header, "multipart/form-data"):
		file, header, err := r.FormFile("file")
		if err != nil {
			render.Render(w, r, response.InternalResponse(err))
			return nil, err
		}
		defer file.Close()

		mimeType := header.Header.Get("Content-Type")

		fileContent, err := io.ReadAll(file)
		if err != nil {
			render.Render(w, r, response.InternalResponse(err))
			return nil, err
		}

		return genai.NewContentFromBytes(fileContent, mimeType, genai.RoleUser), nil

	default:
		err := fmt.Errorf("format %v not accepted", header)
		render.Render(w, r, response.InvalidRequest(err))
		return nil, err
	}
}

func httpHandler(runner *Runner) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sub, err := identity.GetSubFromRequest(r.Context(), r.Header)
		if err != nil {
			render.Render(w, r, response.Unauthorized(fmt.Errorf("no sub found : %w", err)))
			return
		}

		input, err := parseRequest(w, r)
		if err != nil {
			return
		}

		resp, err := runner.Run(RunParams{
			Input:        input,
			UserId:       sub,
			ClearSession: true,
		})
		if err != nil {
			render.Render(w, r, response.InternalResponse(fmt.Errorf("issue with runner %v", err)))
			return
		}

		render.Render(w, r, response.Success(resp))
	}
}
