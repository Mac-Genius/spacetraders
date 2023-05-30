package rest_test

import (
	"path/filepath"
	"testing"

	"github.com/mac-genius/spacetraders/api"
	"github.com/mac-genius/spacetraders/pkg/rest"
)

type restContractTable struct {
	name        string
	input       string
	statusCode  int
	err         string
	contractID  string
	shipSymbol  string
	tradeSymbol string
	units       int32
	expectedLen int
	limit       int
	page        int
	isNil       bool
}

func (a restContractTable) Error() string {
	return a.err
}

func (a restContractTable) InputFile() string {
	return a.input
}

func (a restContractTable) StatusCode() int {
	return a.statusCode
}

func (a restContractTable) ExpectedLen() int {
	return a.expectedLen
}

func (a restContractTable) IsNil() bool {
	return a.isNil
}

func TestAcceptContract(t *testing.T) {
	tables := []restContractTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "accept_contract_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.AcceptContractResponse_Data, error) {
			return rest.AcceptContract(sc, table.contractID)
		}))
	}
}

func TestDeliverContract(t *testing.T) {
	tables := []restContractTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "deliver_contract_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.DeliverContractResponse_Data, error) {
			return rest.DeliverContract(sc, table.contractID, table.shipSymbol, table.tradeSymbol, table.units)
		}))
	}
}

func TestFulfillContract(t *testing.T) {
	tables := []restContractTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "fulfill_contract_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.FulFillContractResponse_Data, error) {
			return rest.FulfillContract(sc, table.contractID)
		}))
	}
}

func TestGetContract(t *testing.T) {
	tables := []restContractTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "get_contract_200.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.Contract, error) {
			return rest.GetContract(sc, table.contractID)
		}))
	}
}

func TestListContracts(t *testing.T) {
	tables := []restContractTable{
		{
			name:        "Success",
			input:       filepath.Join("testdata", "responses", "list_contracts_200.json"),
			statusCode:  200,
			expectedLen: 1,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFuncList(t, table, func(sc *rest.Client) ([]*api.Contract, error) {
			return rest.ListContracts(sc)
		}))
	}
}

func TestListContractsPaging(t *testing.T) {
	tables := []restContractTable{
		{
			name:        "Success",
			input:       filepath.Join("testdata", "responses", "list_contracts_200.json"),
			statusCode:  200,
			expectedLen: 1,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFuncList(t, table, func(sc *rest.Client) ([]*api.Contract, error) {
			return rest.ListContractsPaging(sc, table.limit, table.page)
		}))
	}
}
