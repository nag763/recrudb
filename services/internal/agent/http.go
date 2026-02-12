package agent

import (
	"errors"
	"fmt"
	"nag763/recrudb/api/internal/agent/firestore"
	"nag763/recrudb/api/internal/profile"
	"nag763/recrudb/api/internal/request"
	"nag763/recrudb/api/internal/response"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"google.golang.org/adk/session"
)

func interactionHttpHandler(runner *Runner, profileRepository *profile.FirestoreProfileRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		ctx := r.Context()
		sub := ctx.Value("sub").(string)

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

		state := make(map[string]any)
		if offerId := httpRequest.Context.OfferId; offerId != "" {
			offer, err := profileRepository.GetOffer(ctx, sub, offerId)
			if err != nil {
				render.Render(w, r, response.InvalidRequest(err))
				return
			}
			state["offer"] = offer
		}
		if user := httpRequest.Context.User; user {
			user, err := profileRepository.GetProfile(ctx, sub)
			if err != nil {
				render.Render(w, r, response.InvalidRequest(err))
				return
			}
			state["user"] = user
		}

		resp, err := runner.Run(RunParams{
			Input:        httpRequest.ToGenAiContent(),
			UserId:       sub,
			ClearSession: httpRequest.ClearSession,
			SessionId:    httpRequest.SessionId,
			State:        state,
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
		ctx := r.Context()
		sub := ctx.Value("sub").(string)

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

func sessionGetHttpHandler(runner *Runner, sessionService *firestore.FirestoreSessionService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sub := ctx.Value("sub").(string)

		sessionId := r.PathValue("sessionId")
		if sessionId == "" {
			render.Render(w, r, response.InvalidRequest(errors.New("session id can't be empty")))
			return
		}

		session, err := sessionService.Get(r.Context(), &session.GetRequest{
			AppName:   runner.appName,
			UserID:    sub,
			SessionID: sessionId,
		})
		if err != nil {
			render.Render(w, r, response.InternalResponse(err))
		}

		render.Render(w, r, response.Success(session))
	}
}

func sessionDeleteHttpHandler(runner *Runner, sessionService *firestore.FirestoreSessionService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sub := ctx.Value("sub").(string)

		sessionId := r.PathValue("sessionId")
		if sessionId == "" {
			render.Render(w, r, response.InvalidRequest(errors.New("session id can't be empty")))
			return
		}

		if err := sessionService.Delete(r.Context(), &session.DeleteRequest{
			AppName:   runner.appName,
			UserID:    sub,
			SessionID: sessionId,
		}); err != nil {
			render.Render(w, r, response.InternalResponse(err))
		}

		render.Render(w, r, response.Success("DELETED"))
	}
}
