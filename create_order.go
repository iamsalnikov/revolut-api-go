package revolut_api_go

import (
	"encoding/json"
	"net/http"
)

func (c *Client) CreateOrder(
	order CreateOrderPayload,
) (CreateOrderResponse, error) {
	body, err := json.Marshal(order)
	if err != nil {
		return CreateOrderResponse{}, err
	}

	res, err := c.makeRequest([]string{"orders"}, http.MethodPost, body)
	if err != nil {
		return CreateOrderResponse{}, err
	}

	response := new(CreateOrderResponse)
	err = json.Unmarshal(res, response)
	if err != nil {
		return CreateOrderResponse{}, err
	}

	return *response, nil
}
