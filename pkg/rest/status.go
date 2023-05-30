package rest

import (
	"github.com/mac-genius/spacetraders/api"
)

// Status returns the status of the SpaceTraders REST API.
func Status(c *Client) (*api.StatusResponse, error) {
	return DoRequest(c, EndpointStatus, &api.StatusResponse{})
}
