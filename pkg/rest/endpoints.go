package rest

import "net/http"

// Endpoint describes a SpaceTraders v2 REST endpoint.
type Endpoint struct {
	Method string
	Path   string
}

// General endpoints.
var (
	EndpointRegister = Endpoint{http.MethodPost, "/register"}
	EndpointStatus   = Endpoint{http.MethodGet, "/"}
)

// Agent endpoints.
var (
	EndpointMyAgent = Endpoint{http.MethodGet, "/my/agent"}
)

// Contract endpoints.
var (
	EndpointAcceptContract  = Endpoint{http.MethodPost, "/my/contracts/{{.ContractID}}/accept"}
	EndpointDeliverContract = Endpoint{http.MethodPost, "/my/contracts/{{.ContractID}}/deliver"}
	EndpointFulfillContract = Endpoint{http.MethodPost, "/my/contracts/{{.ContractID}}/fulfill"}
	EndpointGetContract     = Endpoint{http.MethodGet, "/my/contracts/{{.ContractID}}"}
	EndpointListContracts   = Endpoint{http.MethodGet, "/my/contracts"}
)

// Faction endpoints.
var (
	EndpointGetFaction   = Endpoint{http.MethodGet, "/factions/{{.Faction}}"}
	EndpointListFactions = Endpoint{http.MethodGet, "/factions"}
)

// Fleet endpoints.
var (
	EndpointCreateChart       = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/chart"}
	EndpointCreateSurvey      = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/survey"}
	EndpointDockShip          = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/dock"}
	EndpointExtractResources  = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/extract"}
	EndpointGetShip           = Endpoint{http.MethodGet, "/my/ships/{{.ShipSymbol}}"}
	EndpointGetShipCargo      = Endpoint{http.MethodGet, "/my/ships/{{.ShipSymbol}}/cargo"}
	EndpointGetShipCooldown   = Endpoint{http.MethodGet, "/my/ships/{{.ShipSymbol}}/cooldown"}
	EndpointGetShipNav        = Endpoint{http.MethodGet, "/my/ships/{{.ShipSymbol}}/nav"}
	EndpointListShips         = Endpoint{http.MethodGet, "/my/ships"}
	EndpointJettison          = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/jettison"}
	EndpointJumpShip          = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/jump"}
	EndpointNavigateShip      = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/navigate"}
	EndpointNegotiateContract = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/negotiate/contract"}
	EndpointOrbitShip         = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/orbit"}
	EndpointPatchShipNav      = Endpoint{http.MethodPatch, "/my/ships/{{.ShipSymbol}}/nav"}
	EndpointPurchaseCargo     = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/purchase"}
	EndpointPurchaseShip      = Endpoint{http.MethodPost, "/my/ships"}
	EndpointRefuelShip        = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/refuel"}
	EndpointScanShips         = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/scan/ships"}
	EndpointScanSystems       = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/scan/systems"}
	EndpointScanWaypoints     = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/scan/waypoints"}
	EndpointSellCargo         = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/sell"}
	EndpointShipRefine        = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/refine"}
	EndpointTransferCargo     = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/transfer"}
	EndpointWarpShip          = Endpoint{http.MethodPost, "/my/ships/{{.ShipSymbol}}/warp"}
)

// Systems endpoints.
var (
	EndpointGetJumpGate   = Endpoint{http.MethodGet, "/systems/{{.SystemSymbol}}/waypoints/{{.WaypointSymbol}}/jump-gate"}
	EndpointGetMarket     = Endpoint{http.MethodGet, "/systems/{{.SystemSymbol}}/waypoints/{{.WaypointSymbol}}/market"}
	EndpointGetShipyard   = Endpoint{http.MethodGet, "/systems/{{.SystemSymbol}}/waypoints/{{.WaypointSymbol}}/shipyard"}
	EndpointGetSystem     = Endpoint{http.MethodGet, "/systems/{{.SystemSymbol}}"}
	EndpointGetWaypoint   = Endpoint{http.MethodGet, "/systems/{{.SystemSymbol}}/waypoints/{{.WaypointSymbol}}"}
	EndpointListSystems   = Endpoint{http.MethodGet, "/systems"}
	EndpointListWaypoints = Endpoint{http.MethodGet, "/systems/{{.SystemSymbol}}/waypoints"}
)
