package rest_test

import (
	"path/filepath"
	"testing"

	"github.com/mac-genius/spacetraders/api"
	"github.com/mac-genius/spacetraders/pkg/rest"
)

type statusTable struct {
	name       string
	input      string
	statusCode int
	err        string
	isNil      bool
}

func (a statusTable) Error() string {
	return a.err
}

func (a statusTable) InputFile() string {
	return a.input
}

func (a statusTable) StatusCode() int {
	return a.statusCode
}

func (a statusTable) IsNil() bool {
	return a.isNil
}

func TestStatus(t *testing.T) {
	tables := []statusTable{
		{
			name:       "",
			input:      filepath.Join("testdata", "responses", "status.json"),
			statusCode: 200,
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.StatusResponse, error) {
			return rest.Status(sc)
		}))
	}
}
