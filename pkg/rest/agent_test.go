package rest_test

import (
	"path/filepath"
	"testing"

	"github.com/mac-genius/spacetraders/api"
	"github.com/mac-genius/spacetraders/pkg/rest"
)

type agentTable struct {
	name       string
	input      string
	statusCode int
	err        string
	isNil      bool
}

func (a agentTable) Error() string {
	return a.err
}

func (a agentTable) InputFile() string {
	return a.input
}

func (a agentTable) StatusCode() int {
	return a.statusCode
}

func (a agentTable) IsNil() bool {
	return a.isNil
}

func TestMyAgent(t *testing.T) {
	tables := []agentTable{
		{
			name:       "Success",
			input:      filepath.Join("testdata", "responses", "my_agent_200.json"),
			statusCode: 200,
		},
		{
			name:       "Missing Token",
			input:      filepath.Join("testdata", "responses", "my_agent_401.json"),
			statusCode: 401,
			err:        "status 401: Failed to parse token. Token is missing or empty. (4100)",
		},
	}
	for _, table := range tables {
		t.Run(table.name, CreateTestFunc(t, table, func(sc *rest.Client) (*api.Agent, error) {
			return rest.MyAgent(sc)
		}))
	}
}
