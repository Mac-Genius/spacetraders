package rest_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/mac-genius/spacetraders/pkg/rest"
)

type TestTransport struct {
	StatusCode int
	InputFile  string
}

func (t *TestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	data, err := os.ReadFile(t.InputFile)
	if err != nil {
		return nil, err
	}
	return &http.Response{
		Body:       io.NopCloser(bytes.NewBuffer(data)),
		StatusCode: t.StatusCode,
	}, nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(t *TestTransport) *http.Client {
	return &http.Client{
		Transport: t,
	}
}

type restTestTable interface {
	Error() string
	InputFile() string
	StatusCode() int
	IsNil() bool
}

type restTestTableList interface {
	restTestTable
	ExpectedLen() int
}

type restTestCallback[T any] func(*rest.Client) (T, error)

func CreateTestFunc[T comparable](t *testing.T, table restTestTable, callback restTestCallback[T]) func(*testing.T) {
	return func(t *testing.T) {
		var zeroObj T
		client := rest.NewClient(
			NewTestClient(&TestTransport{
				InputFile:  table.InputFile(),
				StatusCode: table.StatusCode(),
			}),
			context.Background(),
			"",
		)
		obj, err := callback(client)
		if table.Error() != "" {
			if err != nil {
				if err.Error() != table.Error() {
					t.Errorf("Got:\n\t%s\nWanted:\n\t%s", err.Error(), table.Error())
				}
			} else {
				t.Error("was expecting an error")
			}
		} else {
			if err != nil {
				t.Fatal(err)
			} else {
				if obj != zeroObj && table.IsNil() {
					t.Error("object should be nil")
				} else if obj == zeroObj && !table.IsNil() {
					t.Error("object should not be nil")
				}
			}
		}
	}
}

type ComparableSlice[T comparable] interface {
	comparable
	~[]T
}

func CreateTestFuncList[T comparable](t *testing.T, table restTestTableList, callback restTestCallback[[]T]) func(*testing.T) {
	return func(t *testing.T) {
		client := rest.NewClient(
			NewTestClient(&TestTransport{
				InputFile:  table.InputFile(),
				StatusCode: table.StatusCode(),
			}),
			context.Background(),
			"",
		)
		obj, err := callback(client)
		if table.Error() != "" {
			if err != nil {
				if err.Error() != table.Error() {
					t.Errorf("Got:\n\t%s\nWanted:\n\t%s", err.Error(), table.Error())
				}
			} else {
				t.Error("was expecting an error")
			}
		} else {
			if err != nil {
				t.Fatal(err)
			} else {
				if len(obj) != table.ExpectedLen() {
					t.Errorf("slices are not the same size\nGot: %d\nWant: %d", len(obj), table.ExpectedLen())
				}
			}
		}
	}
}
