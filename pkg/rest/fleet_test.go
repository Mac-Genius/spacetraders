package rest_test

import (
	"path/filepath"
	"testing"

	"github.com/mac-genius/spacetraders/api"
	"github.com/mac-genius/spacetraders/pkg/rest"
)

type restFleetTable struct {
	name           string
	input          string
	statusCode     int
	err            string
	shipSymbol     string
	expectedLen    int
	symbol         string
	systemSymbol   string
	waypointSymbol string
	units          int32
	flightMode     api.ShipNavFlightMode
	shipType       api.ShipType
	tradeSymbol    string
	targetShip     string
	isNil          bool
	limit          int
	page           int
}

func (a restFleetTable) Error() string {
	return a.err
}

func (a restFleetTable) InputFile() string {
	return a.input
}

func (a restFleetTable) StatusCode() int {
	return a.statusCode
}

func (a restFleetTable) ExpectedLen() int {
	return a.expectedLen
}

func (a restFleetTable) IsNil() bool {
	return a.isNil
}

func TestCreateChart(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "create_chart_201.json"),
			statusCode: 201,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.CreateChartResponse_Data, error) {
			return rest.CreateChart(sc, table.shipSymbol)
		}))
	}
}

func TestCreateSurvey(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "create_survey_201.json"),
			statusCode: 201,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.CreateSurveyResponse_Data, error) {
			return rest.CreateSurvey(sc, table.shipSymbol)
		}))
	}
}

func TestDockShip(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "dock_ship_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.DockShipResponse_Data, error) {
			return rest.DockShip(sc, table.shipSymbol)
		}))
	}
}

func TestExtractResources(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "extract_resources_201.json"),
			statusCode: 201,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.ExtractResourcesResponse_Data, error) {
			return rest.ExtractResources(sc, table.shipSymbol, &api.Survey{})
		}))
	}
}

func TestGetShip(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "get_ship_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.Ship, error) {
			return rest.GetShip(sc, table.shipSymbol)
		}))
	}
}

func TestGetShipCargo(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "get_ship_cargo_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.ShipCargo, error) {
			return rest.GetShipCargo(sc, table.shipSymbol)
		}))
	}
}

func TestGetShipCooldown(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "get_ship_cooldown_200.json"),
			statusCode: 200,
		},
		{
			name:       "Success 204",
			input:      filepath.Join("testdata", "responses", "get_ship_cooldown_204.json"),
			statusCode: 204,
			isNil:      true,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.Cooldown, error) {
			return rest.GetShipCooldown(sc, table.shipSymbol)
		}))
	}
}

func TestGetShipNav(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "get_ship_nav_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.ShipNav, error) {
			return rest.GetShipNav(sc, table.shipSymbol)
		}))
	}
}

func TestJettisonCargo(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "jettison_cargo_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.JettisonCargoResponse_Data, error) {
			return rest.JettisonCargo(sc, table.shipSymbol, table.symbol, table.units)
		}))
	}
}

func TestJumpShip(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "jump_ship_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.JumpShipResponse_Data, error) {
			return rest.JumpShip(sc, table.shipSymbol, table.systemSymbol)
		}))
	}
}

func TestListShips(t *testing.T) {
	tables := []restFleetTable{
		{
			name:        "Success",
			input:       filepath.Join("testdata", "responses", "list_ships_200.json"),
			statusCode:  200,
			expectedLen: 1,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFuncList(t, table, func(sc *rest.Client) ([]*api.Ship, error) {
			return rest.ListShips(sc)
		}))
	}
}

func TestListShipsPaging(t *testing.T) {
	tables := []restFleetTable{
		{
			name:        "Success",
			input:       filepath.Join("testdata", "responses", "list_ships_200.json"),
			statusCode:  200,
			expectedLen: 1,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFuncList(t, table, func(sc *rest.Client) ([]*api.Ship, error) {
			return rest.ListShipsPaging(sc, table.limit, table.page)
		}))
	}
}

func TestNavigateShip(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "navigate_ship_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.NavigateShipResponse_Data, error) {
			return rest.NavigateShip(sc, table.shipSymbol, table.waypointSymbol)
		}))
	}
}

func TestNegotiateContract(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "negotiate_contract_201.json"),
			statusCode: 201,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.NegotiateContractResponse_Data, error) {
			return rest.NegotiateContract(sc, table.shipSymbol)
		}))
	}
}

func TestOrbitShip(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "orbit_ship_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.OrbitShipResponse_Data, error) {
			return rest.OrbitShip(sc, table.shipSymbol)
		}))
	}
}

func TestPatchShipNav(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "patch_ship_nav_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.ShipNav, error) {
			return rest.PatchShipNav(sc, table.shipSymbol, table.flightMode)
		}))
	}
}

func TestPurchaseCargo(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "purchase_cargo_201.json"),
			statusCode: 201,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.CargoTransactionResponse_Data, error) {
			return rest.PurchaseCargo(sc, table.shipSymbol, table.symbol, table.units)
		}))
	}
}

func TestPurchaseShip(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "purchase_ship_201.json"),
			statusCode: 201,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.PurchaseShipResponse_Data, error) {
			return rest.PurchaseShip(sc, table.shipType, table.waypointSymbol)
		}))
	}
}

func TestRefuelShip(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "refuel_ship_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.RefuelShipResponse_Data, error) {
			return rest.RefuelShip(sc, table.shipSymbol)
		}))
	}
}

func TestScanShips(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "scan_ships_201.json"),
			statusCode: 201,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.ScanShipsResponse_Data, error) {
			return rest.ScanShips(sc, table.shipSymbol)
		}))
	}
}

func TestScanSystems(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "scan_systems_201.json"),
			statusCode: 201,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.ScanSystemsResponse_Data, error) {
			return rest.ScanSystems(sc, table.shipSymbol)
		}))
	}
}

func TestScanWaypoints(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "scan_waypoints_201.json"),
			statusCode: 201,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.ScanWaypointsResponse_Data, error) {
			return rest.ScanWaypoints(sc, table.shipSymbol)
		}))
	}
}

func TestSellCargo(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "sell_cargo_201.json"),
			statusCode: 201,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.CargoTransactionResponse_Data, error) {
			return rest.SellCargo(sc, table.shipSymbol, table.symbol, table.units)
		}))
	}
}

func TestShipRefine(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "ship_refine_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.ShipRefineResponse_Data, error) {
			return rest.ShipRefine(sc, table.shipSymbol)
		}))
	}
}

func TestTransferCargo(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "transfer_cargo_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.TransferCargoResponse_Data, error) {
			return rest.TransferCargo(sc, table.shipSymbol, table.tradeSymbol, table.units, table.targetShip)
		}))
	}
}

func TestWarpShip(t *testing.T) {
	tables := []restFleetTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "warp_ship_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.WarpShipResponse_Data, error) {
			return rest.WarpShip(sc, table.shipSymbol, table.waypointSymbol)
		}))
	}
}
