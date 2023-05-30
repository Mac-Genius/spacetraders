package rest_test

import (
	"path/filepath"
	"testing"

	"github.com/mac-genius/spacetraders/api"
	"github.com/mac-genius/spacetraders/pkg/rest"
)

type restSystemsTable struct {
	name           string
	input          string
	statusCode     int
	err            string
	systemSymbol   string
	waypointSymbol string
	expectedLen    int
	limit          int
	page           int
	isNil          bool
}

func (a restSystemsTable) Error() string {
	return a.err
}

func (a restSystemsTable) InputFile() string {
	return a.input
}

func (a restSystemsTable) StatusCode() int {
	return a.statusCode
}

func (a restSystemsTable) ExpectedLen() int {
	return a.expectedLen
}

func (a restSystemsTable) IsNil() bool {
	return a.isNil
}

func TestGetJumpGate(t *testing.T) {
	tables := []restSystemsTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "get_jumpgate_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.JumpGate, error) {
			return rest.GetJumpGate(sc, table.systemSymbol, table.waypointSymbol)
		}))
	}
}

func TestGetMarket(t *testing.T) {
	tables := []restSystemsTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "get_market_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.Market, error) {
			return rest.GetMarket(sc, table.systemSymbol, table.waypointSymbol)
		}))
	}
}

func TestGetShipyard(t *testing.T) {
	tables := []restSystemsTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "get_shipyard_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.Shipyard, error) {
			return rest.GetShipyard(sc, table.systemSymbol, table.waypointSymbol)
		}))
	}
}

func TestGetSystem(t *testing.T) {
	tables := []restSystemsTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "get_system_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.System, error) {
			return rest.GetSystem(sc, table.systemSymbol)
		}))
	}
}

func TestGetWaypoint(t *testing.T) {
	tables := []restSystemsTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "get_waypoint_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.Waypoint, error) {
			return rest.GetWaypoint(sc, table.systemSymbol, table.waypointSymbol)
		}))
	}
}

func TestListSystems(t *testing.T) {
	tables := []restSystemsTable{
		{
			name:        "Success",
			input:       filepath.Join("testdata", "responses", "list_systems_200.json"),
			statusCode:  200,
			expectedLen: 1,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFuncList(t, table, func(sc *rest.Client) ([]*api.System, error) {
			return rest.ListSystems(sc)
		}))
	}
}

func TestListSystemsPaging(t *testing.T) {
	tables := []restSystemsTable{
		{
			name:        "Success",
			input:       filepath.Join("testdata", "responses", "list_systems_200.json"),
			statusCode:  200,
			expectedLen: 1,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFuncList(t, table, func(sc *rest.Client) ([]*api.System, error) {
			return rest.ListSystemsPaging(sc, table.limit, table.page)
		}))
	}
}

func TestListWaypoints(t *testing.T) {
	tables := []restSystemsTable{
		{
			name:        "Success",
			input:       filepath.Join("testdata", "responses", "list_waypoints_200.json"),
			statusCode:  200,
			expectedLen: 1,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFuncList(t, table, func(sc *rest.Client) ([]*api.Waypoint, error) {
			return rest.ListWaypoints(sc, table.systemSymbol)
		}))
	}
}

func TestListWaypointsPaging(t *testing.T) {
	tables := []restSystemsTable{
		{
			name:        "Success",
			input:       filepath.Join("testdata", "responses", "list_waypoints_200.json"),
			statusCode:  200,
			expectedLen: 1,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFuncList(t, table, func(sc *rest.Client) ([]*api.Waypoint, error) {
			return rest.ListWaypointsPaging(sc, table.systemSymbol, table.limit, table.page)
		}))
	}
}
