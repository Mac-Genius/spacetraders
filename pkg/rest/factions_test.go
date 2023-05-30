package rest_test

import (
	"path/filepath"
	"testing"

	"github.com/mac-genius/spacetraders/api"
	"github.com/mac-genius/spacetraders/pkg/rest"
)

type restFactionTable struct {
	name        string
	input       string
	statusCode  int
	err         string
	faction     string
	expectedLen int
	limit       int
	page        int
	isNil       bool
}

func (a restFactionTable) Error() string {
	return a.err
}

func (a restFactionTable) InputFile() string {
	return a.input
}

func (a restFactionTable) StatusCode() int {
	return a.statusCode
}

func (a restFactionTable) ExpectedLen() int {
	return a.expectedLen
}

func (a restFactionTable) IsNil() bool {
	return a.isNil
}

func TestGetFaction(t *testing.T) {
	tables := []restFactionTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "get_faction_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.Faction, error) {
			return rest.GetFaction(sc, table.faction)
		}))
	}
}

func TestListFactions(t *testing.T) {
	tables := []restFactionTable{
		{
			name:        "Success",
			input:       filepath.Join("testdata", "responses", "list_factions_200.json"),
			statusCode:  200,
			expectedLen: 1,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFuncList(t, table, func(sc *rest.Client) ([]*api.Faction, error) {
			return rest.ListFactions(sc)
		}))
	}
}

func TestListFactionsPaging(t *testing.T) {
	tables := []restFactionTable{
		{
			name:        "Success",
			input:       filepath.Join("testdata", "responses", "list_factions_200.json"),
			statusCode:  200,
			expectedLen: 1,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFuncList(t, table, func(sc *rest.Client) ([]*api.Faction, error) {
			return rest.ListFactionsPaging(sc, table.limit, table.page)
		}))
	}
}
