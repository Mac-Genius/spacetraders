package rest

import "github.com/mac-genius/spacetraders/api"

type shipQuery struct {
	ShipSymbol string
}

// CreateChart commands a ship to chart the current waypoint.
//
// Waypoints in the universe are uncharted by default. These locations will not
// show up in the API until they have been charted by a ship.
//
// Charting a location will record your agent as the one who created the chart.
func CreateChart(c *Client, shipSymbol string) (*api.CreateChartResponse_Data, error) {
	return unwrapData[*api.CreateChartResponse_Data](DoRequestFull(c, EndpointCreateChart, &api.CreateChartResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// CreateSurvey surveys the current waypoint. If you want to target specific
// yields for an extraction, you can survey a waypoint, such as an asteroid field,
// and send the survey in the body of the extract request. Each survey may have
// multiple deposits, and if a symbol shows up more than once, that indicates a
// higher chance of extracting that resource.
//
// Your ship will enter a cooldown between consecutive survey requests. Surveys
// will eventually expire after a period of time. Multiple ships can use the same
// survey for extraction.
func CreateSurvey(c *Client, shipSymbol string) (*api.CreateSurveyResponse_Data, error) {
	return unwrapData[*api.CreateSurveyResponse_Data](DoRequestFull(c, EndpointCreateSurvey, &api.CreateSurveyResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// DockShip attempts to dock your ship at it's current location. Docking will
// only succeed if the waypoint is a dockable location, and your ship is
// capable of docking at the time of the request.
func DockShip(c *Client, shipSymbol string) (*api.DockShipResponse_Data, error) {
	return unwrapData[*api.DockShipResponse_Data](DoRequestFull(c, EndpointDockShip, &api.DockShipResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// ExtractResources extracts resources from the waypoint into your ship. Send
// an optional survey as the payload to target specific yields.
func ExtractResources(c *Client, shipSymbol string, survey *api.Survey) (*api.ExtractResourcesResponse_Data, error) {
	return unwrapData[*api.ExtractResourcesResponse_Data](DoRequestFull(c, EndpointDockShip, &api.ExtractResourcesResponse{}, survey, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// GetShip retrieves the details of your ship.
func GetShip(c *Client, shipSymbol string) (*api.Ship, error) {
	return unwrapData[*api.Ship](DoRequestFull(c, EndpointGetShip, &api.GetShipResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// GetShipCargo retrieves the cargo of your ship.
func GetShipCargo(c *Client, shipSymbol string) (*api.ShipCargo, error) {
	return unwrapData[*api.ShipCargo](DoRequestFull(c, EndpointGetShipCargo, &api.GetShipCargoResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// GetShipCooldown retrieves the details of your ship's reactor cooldown. Some
// actions such as activating your jump drive, scanning, or extracting resources
// taxes your reactor and results in a cooldown.
//
// Your ship cannot perform additional actions until your cooldown has expired.
// The duration of your cooldown is relative to the power consumption of the
// related modules or mounts for the action taken.
func GetShipCooldown(c *Client, shipSymbol string) (*api.Cooldown, error) {
	return unwrapData[*api.Cooldown](DoRequestFull(c, EndpointGetShipCooldown, &api.GetShipCooldownResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// GetShipNav gets the current nav status of a ship.
func GetShipNav(c *Client, shipSymbol string) (*api.ShipNav, error) {
	return unwrapData[*api.ShipNav](DoRequestFull(c, EndpointGetShipNav, &api.GetShipNavResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// JettisonCargo jettisons cargo from your ship's cargo hold.
func JettisonCargo(c *Client, shipSymbol string, symbol string, units int32) (*api.JettisonCargoResponse_Data, error) {
	return unwrapData[*api.JettisonCargoResponse_Data](DoRequestFull(c, EndpointJettison, &api.JettisonCargoResponse{}, &api.JettisonCargoRequest{
		Symbol: symbol,
		Units:  units,
	}, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// JumpShip jumps your ship instantly to a target system. When used while in
// orbit or docked to a jump gate waypoint, any ship can use this command. When
// used elsewhere, jumping requires a jump drive unit and consumes a unit of
// antimatter (which needs to be in your cargo).
func JumpShip(c *Client, shipSymbol string, systemSymbol string) (*api.JumpShipResponse_Data, error) {
	return unwrapData[*api.JumpShipResponse_Data](DoRequestFull(c, EndpointJumpShip, &api.JumpShipResponse{}, &api.JumpShipRequest{
		SystemSymbol: systemSymbol,
	}, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// ListShips retrieves all of your ships.
func ListShips(c *Client) ([]*api.Ship, error) {
	return DoPageRequest[*api.Ship](c, EndpointListShips, &api.ListShipsResponse{}, MaxPageSize)
}

// ListShips retrieves all of your ships of size `limit` on a given page.
func ListShipsPaging(c *Client, limit int, page int) ([]*api.Ship, error) {
	return unwrapData[[]*api.Ship](DoRequestFull(c, EndpointListShips, &api.ListShipsResponse{},
		nil, nil, map[string]interface{}{
			"limit": limit,
			"page":  page,
		}))
}

// NavigateShip navigates to a target destination. The destination must be
// located within the same system as the ship. Navigating will consume the
// necessary fuel and supplies from the ship's manifest, and will pay out crew
// wages from the agent's account.
//
// The returned response will detail the route information including the
// expected time of arrival. Most ship actions are unavailable until the ship has
// arrived at it's destination.
//
// To travel between systems, see the ship's warp or jump actions.
func NavigateShip(c *Client, shipSymbol string, waypointSymbol string) (*api.NavigateShipResponse_Data, error) {
	return unwrapData[*api.NavigateShipResponse_Data](DoRequestFull(c, EndpointNavigateShip, &api.NavigateShipResponse{}, &api.NavigateShipRequest{
		WaypointSymbol: waypointSymbol,
	}, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// NegotiateContract generates a new contract.
func NegotiateContract(c *Client, shipSymbol string) (*api.NegotiateContractResponse_Data, error) {
	return unwrapData[*api.NegotiateContractResponse_Data](DoRequestFull(c, EndpointNegotiateContract, &api.NegotiateContractResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// OrbitShip attempts to move your ship into orbit at it's current location. The
// request will only succeed if your ship is capable of moving into orbit at the
// time of the request.
func OrbitShip(c *Client, shipSymbol string) (*api.OrbitShipResponse_Data, error) {
	return unwrapData[*api.OrbitShipResponse_Data](DoRequestFull(c, EndpointOrbitShip, &api.OrbitShipResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// PatchShipNav updates the nav data of a ship, such as the flight mode.
func PatchShipNav(c *Client, shipSymbol string, flightMode api.ShipNavFlightMode) (*api.ShipNav, error) {
	return unwrapData[*api.ShipNav](DoRequestFull(c, EndpointPatchShipNav, &api.PatchShipNavResponse{}, &api.PatchShipNavRequest{
		FlightMode: flightMode,
	}, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// PurchaseCargo purchases cargo.
func PurchaseCargo(c *Client, shipSymbol string, symbol string, units int32) (*api.CargoTransactionResponse_Data, error) {
	return unwrapData[*api.CargoTransactionResponse_Data](DoRequestFull(c, EndpointPurchaseCargo, &api.CargoTransactionResponse{}, &api.CargoTransactionRequest{
		Symbol: symbol,
		Units:  units,
	}, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// PurchaseShip purchases a ship.
func PurchaseShip(c *Client, shipType api.ShipType, waypointSymbol string) (*api.PurchaseShipResponse_Data, error) {
	return unwrapData[*api.PurchaseShipResponse_Data](DoRequestFull(c, EndpointPurchaseShip, &api.PurchaseShipResponse{},
		&api.PurchaseShipRequest{
			ShipType:       shipType,
			WaypointSymbol: waypointSymbol,
		}, nil, nil))
}

// RefuelShip refuel your ship from the local market.
func RefuelShip(c *Client, shipSymbol string) (*api.RefuelShipResponse_Data, error) {
	return unwrapData[*api.RefuelShipResponse_Data](DoRequestFull(c, EndpointRefuelShip, &api.RefuelShipResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// ScanShips activates your ship's sensor arrays to scan for ship information.
func ScanShips(c *Client, shipSymbol string) (*api.ScanShipsResponse_Data, error) {
	return unwrapData[*api.ScanShipsResponse_Data](DoRequestFull(c, EndpointScanShips, &api.ScanShipsResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// ScanSystems activates your ship's sensor arrays to scan for system information.
func ScanSystems(c *Client, shipSymbol string) (*api.ScanSystemsResponse_Data, error) {
	return unwrapData[*api.ScanSystemsResponse_Data](DoRequestFull(c, EndpointScanSystems, &api.ScanSystemsResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// ScanWaypoints activates your ship's sensor arrays to scan for waypoint information.
func ScanWaypoints(c *Client, shipSymbol string) (*api.ScanWaypointsResponse_Data, error) {
	return unwrapData[*api.ScanWaypointsResponse_Data](DoRequestFull(c, EndpointScanWaypoints, &api.ScanWaypointsResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// SellCargo sells cargo.
func SellCargo(c *Client, shipSymbol string, symbol string, units int32) (*api.CargoTransactionResponse_Data, error) {
	return unwrapData[*api.CargoTransactionResponse_Data](DoRequestFull(c, EndpointSellCargo, &api.CargoTransactionResponse{}, &api.CargoTransactionRequest{
		Symbol: symbol,
		Units:  units,
	}, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// ShipRefine attempts to refine the raw materials on your ship. The request
// will only succeed if your ship is capable of refining at the time of the
// request.
func ShipRefine(c *Client, shipSymbol string) (*api.ShipRefineResponse_Data, error) {
	return unwrapData[*api.ShipRefineResponse_Data](DoRequestFull(c, EndpointShipRefine, &api.ShipRefineResponse{}, nil, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// TransferCargo transfers cargo between ships.
func TransferCargo(c *Client, shipSymbol string, tradeSymbol string, units int32, targetShip string) (*api.TransferCargoResponse_Data, error) {
	return unwrapData[*api.TransferCargoResponse_Data](DoRequestFull(c, EndpointTransferCargo, &api.TransferCargoResponse{}, &api.TransferCargoRequest{
		TradeSymbol: tradeSymbol,
		Units:       units,
		ShipSymbol:  targetShip,
	}, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}

// WarpShip warps your ship to a target destination in another system. Warping
// will consume the necessary fuel and supplies from the ship's manifest, and will
// pay out crew wages from the agent's account.

// The returned response will detail the route information including the
// expected time of arrival. Most ship actions are unavailable until the ship
// has arrived at it's destination.
func WarpShip(c *Client, shipSymbol string, waypointSymbol string) (*api.WarpShipResponse_Data, error) {
	return unwrapData[*api.WarpShipResponse_Data](DoRequestFull(c, EndpointWarpShip, &api.WarpShipResponse{}, &api.WarpShipRequest{
		WaypointSymbol: waypointSymbol,
	}, shipQuery{
		ShipSymbol: shipSymbol,
	}, nil))
}
