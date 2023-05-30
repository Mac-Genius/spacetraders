package rest

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mac-genius/spacetraders/api"
	"golang.org/x/time/rate"
	"google.golang.org/protobuf/encoding/protojson"
)

// SpaceTradersV2_URL is the default URL for the SpaceTraders REST API.
const SpaceTradersV2_URL = "https://api.spacetraders.io/v2"

// SpaceTradersV2_MockURL is the default URL for the SpaceTraders mock REST API.
const SpaceTradersV2_MockURL = "https://stoplight.io/mocks/spacetraders/spacetraders/96627693"

// DefaultClient is the default client without authentication.
var DefaultClient = &Client{
	ctx:         context.Background(),
	client:      http.DefaultClient,
	RateLimiter: rate.NewLimiter(rate.Every(time.Second), 2),
	Url:         SpaceTradersV2_URL,
}

// Client is used to interact with the SpaceTraders REST API.
type Client struct {
	client      *http.Client
	ctx         context.Context
	RateLimiter *rate.Limiter
	Url         string
}

// NewClient creates a new client with a given HTTP client, context and URL.
func NewClient(httpClient *http.Client, ctx context.Context, url string) *Client {
	return &Client{
		ctx:         ctx,
		client:      httpClient,
		RateLimiter: rate.NewLimiter(rate.Every(time.Second), 2),
		Url:         url,
	}
}

// NewClientWithAuth creates a new client with the given authentication token.
func NewClientWithAuth(token string) *Client {
	return &Client{
		ctx:         context.WithValue(context.Background(), agentToken, token),
		client:      http.DefaultClient,
		RateLimiter: rate.NewLimiter(rate.Every(time.Second), 2),
		Url:         SpaceTradersV2_URL,
	}
}

// Do calls the underlying HTTP client with the given request.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	if c.RateLimiter != nil {
		err := c.RateLimiter.Wait(c.ctx)
		if err != nil {
			return nil, err
		}
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if val := c.ctx.Value(agentToken); val != nil {
		req.Header.Add("Autorization", "Bearer "+val.(string))
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		respErr := &api.ErrorResponse{}
		data, err := io.ReadAll(resp.Body)
		if err == nil {
			err = protojson.Unmarshal(data, respErr)
			if err == nil {
				return nil, fmt.Errorf("status %d: %s (%d)", resp.StatusCode, respErr.Error.Message, respErr.Error.Code)
			} else {
				return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(data))
			}
		}
		return nil, fmt.Errorf("received %d status code", resp.StatusCode)
	}
	return resp, nil
}
