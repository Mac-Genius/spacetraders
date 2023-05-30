package rest

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"text/template"

	"github.com/mac-genius/spacetraders/api"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// MaxPageSize is the maximum size of a page allowed by the SpaceTraders API.
const MaxPageSize = 20

// DoRequest is a simplified way of making a REST request to the SpaceTraders API.
func DoRequest[T proto.Message](c *Client, e Endpoint, t T) (T, error) {
	return DoRequestFull(c, e, t, nil, nil, nil)
}

// DoRequestWithBody is a simplified way of making a REST request with data to the SpaceTraders API.
func DoRequestWithBody[T proto.Message](c *Client, e Endpoint, t T, body proto.Message) (T, error) {
	return DoRequestFull(c, e, t, body, nil, nil)
}

// DoRequestFull makes a REST request to the SpaceTraders API.
func DoRequestFull[T proto.Message](c *Client, e Endpoint, t T, body proto.Message,
	pathParam interface{}, queryParams map[string]interface{}) (T, error) {
	var rtnVal T
	var reader io.Reader
	if body != nil {
		data, err := protojson.MarshalOptions{
			EmitUnpopulated: true,
		}.Marshal(body)
		if err != nil {
			return rtnVal, err
		}
		fmt.Println(string(data))
		reader = bytes.NewBuffer(data)
	}
	restUrl := c.Url + e.Path
	if pathParam != nil {
		t, err := template.New("_0").Parse(restUrl)
		if err != nil {
			return rtnVal, err
		}
		output := &bytes.Buffer{}
		err = t.Execute(output, pathParam)
		if err != nil {
			return rtnVal, err
		}
		restUrl = output.String()
	}
	if queryParams != nil {
		i := 0
		for key, val := range queryParams {
			if i == 0 {
				restUrl += fmt.Sprintf("?%s=%s", url.QueryEscape(key), url.QueryEscape(fmt.Sprint(val)))
			} else {
				restUrl += fmt.Sprintf("&%s=%s", url.QueryEscape(key), url.QueryEscape(fmt.Sprint(val)))
			}
			i++
		}
	}
	req, err := http.NewRequest(e.Method, restUrl, reader)
	if err != nil {
		return rtnVal, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return rtnVal, err
	}
	if resp.StatusCode == 204 {
		return rtnVal, nil
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return rtnVal, err
	}
	err = protojson.Unmarshal(data, t)
	if err != nil {
		return rtnVal, err
	}
	return t, nil
}

// DoPageRequest is a simplified way of making a REST request with paging to the SpaceTraders API.
func DoPageRequest[T proto.Message](c *Client, e Endpoint, k api.DataListWrapper[T], limit int) ([]T, error) {
	return DoPageRequestFull(c, e, k, limit, nil)
}

// DoPageRequestFull makes multiple REST requests to the SpaceTraders API and returns all results.
func DoPageRequestFull[T proto.Message](c *Client, e Endpoint, k api.DataListWrapper[T], limit int, pathParam interface{}) ([]T, error) {
	list := []T{}
	page := 1
	total := -1
	found := 0
	for {
		data, err := DoRequestFull(c, e, k, nil, pathParam, map[string]interface{}{
			"limit": limit,
			"page":  page,
		})
		if err != nil {
			return nil, err
		}
		found += len(data.GetData())
		total = int(data.GetMeta().GetTotal())
		list = append(list, data.GetData()...)
		if found >= total {
			break
		}
		page++
	}
	return list, nil
}

// unwrapData returns the underlying data from a wrapper.
func unwrapData[T any](a api.DataWrapper[T], err error) (T, error) {
	var rtnVal T
	if a == nil {
		return rtnVal, nil
	}
	if err != nil {
		return rtnVal, err
	}
	return a.GetData(), nil
}
