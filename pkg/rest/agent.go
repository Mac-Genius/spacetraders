package rest

import (
	"github.com/mac-genius/spacetraders/api"
)

// agentContext is used for internal context keys.
type agentContext string

const (
	agentToken agentContext = "token" // Authentication token for a single agent.
)

// MyAgent fetches the agent given the current space client.
func MyAgent(c *Client) (*api.Agent, error) {
	return unwrapData[*api.Agent](DoRequest(c, EndpointMyAgent, &api.AgentData{}))
}
