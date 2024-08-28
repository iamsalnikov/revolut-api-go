package revolut_merchant

import (
	"encoding/json"
	"github.com/jerethom/revolut-api-go/merchant/1.0/types/order_types"
	"net/http"
)

func (c *Merchant) CreateOrder(
	order order_types.CreateOrderPayload,
) (order_types.CreateOrderResponse, error) {
	body, err := json.Marshal(order)
	if err != nil {
		return order_types.CreateOrderResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest([]string{"orders"}, http.MethodPost, body)
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
