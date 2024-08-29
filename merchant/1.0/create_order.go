package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) CreateOrder(
	order CreateOrderPayload,
) (CreateOrderResponse, error) {
	body, err := json.Marshal(order)
	if err != nil {
		return CreateOrderResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest([]string{"orders"}, http.MethodPost, body, url.Values{})
	if err != nil {
		return CreateOrderResponse{}, err
	}

	response := CreateOrderResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
