package rest

import "github.com/mac-genius/spacetraders/api"

// AcceptContract accepts a contract.
func AcceptContract(c *Client, contractID string) (*api.AcceptContractResponse_Data, error) {
	return unwrapData[*api.AcceptContractResponse_Data](DoRequestFull(c, EndpointAcceptContract, &api.AcceptContractResponse{}, nil, struct {
		ContractID string
	}{
		ContractID: contractID,
	}, nil))
}

// DeliverContract delivers cargo on a given contract.
func DeliverContract(c *Client, contractID, shipSymbol, tradeSymbol string, units int32) (*api.DeliverContractResponse_Data, error) {
	return unwrapData[*api.DeliverContractResponse_Data](DoRequestFull(c, EndpointDeliverContract, &api.DeliverContractResponse{}, &api.DeliverContractRequest{
		ShipSymbol:  shipSymbol,
		TradeSymbol: tradeSymbol,
		Units:       units,
	}, struct {
		ContractID string
	}{
		ContractID: contractID,
	}, nil))
}

// FulfillContract fulfills a contract.
func FulfillContract(c *Client, contractID string) (*api.FulFillContractResponse_Data, error) {
	return unwrapData[*api.FulFillContractResponse_Data](DoRequestFull(c, EndpointFulfillContract, &api.FulFillContractResponse{}, nil, struct {
		ContractID string
	}{
		ContractID: contractID,
	}, nil))
}

// GetContract retrieves the details of a contract by ID.
func GetContract(c *Client, contractID string) (*api.Contract, error) {
	return unwrapData[*api.Contract](DoRequestFull(c, EndpointGetContract, &api.GetContractResponse{}, nil, struct {
		ContractID string
	}{
		ContractID: contractID,
	}, nil))
}

// ListContracts retrieves all of your contracts.
func ListContracts(c *Client) ([]*api.Contract, error) {
	return DoPageRequest[*api.Contract](c, EndpointListContracts, &api.ListContractsResponse{}, MaxPageSize)
}

// ListContracts retrieves all of your contracts.
func ListContractsPaging(c *Client, limit int, page int) ([]*api.Contract, error) {
	return unwrapData[[]*api.Contract](DoRequestFull(c, EndpointListContracts, &api.ListContractsResponse{},
		nil, nil, map[string]interface{}{
			"limit": limit,
			"page":  page,
		}))
}
