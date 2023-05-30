package rest

import "github.com/mac-genius/spacetraders/api"

type systemQuery struct {
	SystemSymbol   string
	WaypointSymbol string
}

// GetJumpGate retrieves jump gate details for a waypoint.
func GetJumpGate(c *Client, systemSymbol, waypointSymbol string) (*api.JumpGate, error) {
	return unwrapData[*api.JumpGate](DoRequestFull(c, EndpointGetJumpGate, &api.GetJumpGateResponse{}, nil, systemQuery{
		SystemSymbol:   systemSymbol,
		WaypointSymbol: waypointSymbol,
	}, nil))
}

// GetMarket retrieves imports, exports and exchange data from a marketplace.
// Imports can be sold, exports can be purchased, and exchange goods can be
// purchased or sold. Send a ship to the waypoint to access trade good prices
// and recent transactions.
func GetMarket(c *Client, systemSymbol, waypointSymbol string) (*api.Market, error) {
	return unwrapData[*api.Market](DoRequestFull(c, EndpointGetMarket, &api.GetMarketResponse{}, nil, systemQuery{
		SystemSymbol:   systemSymbol,
		WaypointSymbol: waypointSymbol,
	}, nil))
}

// GetShipyard retrieves the shipyard for a waypoint. Send a ship to the
// waypoint to access ships that are currently available for purchase and
// recent transactions.
func GetShipyard(c *Client, systemSymbol, waypointSymbol string) (*api.Shipyard, error) {
	return unwrapData[*api.Shipyard](DoRequestFull(c, EndpointGetShipyard, &api.GetShipyardResponse{}, nil, systemQuery{
		SystemSymbol:   systemSymbol,
		WaypointSymbol: waypointSymbol,
	}, nil))
}

// GetSystem retrieves the details of a system.
func GetSystem(c *Client, systemSymbol string) (*api.System, error) {
	return unwrapData[*api.System](DoRequestFull(c, EndpointGetSystem, &api.GetSystemResponse{}, nil, systemQuery{
		SystemSymbol: systemSymbol,
	}, nil))
}

// GetWaypoint retrieves the details of a waypoint.
func GetWaypoint(c *Client, systemSymbol, waypointSymbol string) (*api.Waypoint, error) {
	return unwrapData[*api.Waypoint](DoRequestFull(c, EndpointGetWaypoint, &api.GetWaypointResponse{}, nil, systemQuery{
		SystemSymbol:   systemSymbol,
		WaypointSymbol: waypointSymbol,
	}, nil))
}

// ListSystems retrieves a list of all systems.
func ListSystems(c *Client) ([]*api.System, error) {
	return DoPageRequest[*api.System](c, EndpointListSystems, &api.ListSystemsResponse{}, MaxPageSize)
}

// ListSystems retrieves a list of systems of size `limit` on a given page.
func ListSystemsPaging(c *Client, limit int, page int) ([]*api.System, error) {
	return unwrapData[[]*api.System](DoRequestFull(c, EndpointListSystems, &api.ListSystemsResponse{},
		nil, nil, map[string]interface{}{
			"limit": limit,
			"page":  page,
		}))
}

// ListWaypoints retrieves a list of all waypoints in a system.
func ListWaypoints(c *Client, systemSymbol string) ([]*api.Waypoint, error) {
	return DoPageRequestFull[*api.Waypoint](c, EndpointListWaypoints, &api.ListWaypointsResponse{}, MaxPageSize, systemQuery{
		SystemSymbol: systemSymbol,
	})
}

// ListWaypoints retrieves a list of waypoints in a system of size `limit` on a given page.
func ListWaypointsPaging(c *Client, systemSymbol string, limit int, page int) ([]*api.Waypoint, error) {
	return unwrapData[[]*api.Waypoint](DoRequestFull(c, EndpointListWaypoints, &api.ListWaypointsResponse{},
		nil, systemQuery{
			SystemSymbol: systemSymbol,
		}, map[string]interface{}{
			"limit": limit,
			"page":  page,
		}))
}
