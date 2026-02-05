package injection

import (
	"google.golang.org/adk/agent"
	"google.golang.org/adk/session"
)

type HttpAgent struct {
	Path    string
	Agent   agent.Agent
	Session session.Service
}
