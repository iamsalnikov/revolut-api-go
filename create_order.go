package revolut

import (
	"encoding/json"
	"github.com/jerethom/revolut-api-go/types/order_types"
	"net/http"
)

func (c *Client) CreateOrder(
	order order_types.CreateOrderPayload,
) (order_types.CreateOrderResponse, error) {
	body, err := json.Marshal(order)
	if err != nil {
		return order_types.CreateOrderResponse{}, err
	}

	res, err := c.makeRequest([]string{"orders"}, http.MethodPost, body)
	if err != nil {
		return order_types.CreateOrderResponse{}, err
	}

	response := new(order_types.CreateOrderResponse)
	err = json.Unmarshal(res, response)
	if err != nil {
		return order_types.CreateOrderResponse{}, err
	}

	return *response, nil
}
