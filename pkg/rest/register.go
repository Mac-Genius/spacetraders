package rest

import (
	"github.com/mac-genius/spacetraders/api"
)

// Register registers a new agent with SpaceTraders.
func Register(c *Client, faction api.RegisterFaction, symbol, email string) (*api.RegisterResponse_Data, error) {
	return unwrapData[*api.RegisterResponse_Data](DoRequestWithBody(c, EndpointRegister,
		&api.RegisterResponse{}, &api.RegisterRequest{
			Faction: faction,
			Symbol:  symbol,
			Email:   email,
		}))
}
