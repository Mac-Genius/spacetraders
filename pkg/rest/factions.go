package rest

import (
	"github.com/mac-genius/spacetraders/api"
)

// GetFaction returns a single faction for the given name.
func GetFaction(c *Client, name string) (*api.Faction, error) {
	return unwrapData[*api.Faction](DoRequestFull(c, EndpointGetFaction, &api.GetFactionResponse{}, nil, struct {
		Faction string
	}{
		Faction: name,
	}, nil))
}

// ListFactions returns all the factions.
func ListFactions(c *Client) ([]*api.Faction, error) {
	return DoPageRequest[*api.Faction](c, EndpointListFactions, &api.ListFactionsResponse{}, MaxPageSize)
}

// ListFactionsPaging returns a list of factions of size `limit` on a given page.
func ListFactionsPaging(c *Client, limit int, page int) ([]*api.Faction, error) {
	return unwrapData[[]*api.Faction](DoRequestFull(c, EndpointListFactions, &api.ListFactionsResponse{},
		nil, nil, map[string]interface{}{
			"limit": limit,
			"page":  page,
		}))
}
