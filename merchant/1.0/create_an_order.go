package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) CreateAnOrder(
	order CreateAnOrderPayload,
) (CreateAnOrderResponse, error) {
	body, err := json.Marshal(order)
	if err != nil {
		return CreateAnOrderResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest([]string{"orders"}, http.MethodPost, body, url.Values{})
	if err != nil {
		return CreateAnOrderResponse{}, err
	}

	response := CreateAnOrderResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
