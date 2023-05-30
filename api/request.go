package api

import "encoding/json"

type SpaceTraderMetadata struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type SpaceTraderData struct {
	Data json.RawMessage      `json:"data"`
	Meta *SpaceTraderMetadata `json:"meta,omitempty"`
}
