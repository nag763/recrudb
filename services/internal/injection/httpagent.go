package injection

import "google.golang.org/adk/agent"

type HttpAgent struct {
	Path  string
	Agent agent.Agent
}
