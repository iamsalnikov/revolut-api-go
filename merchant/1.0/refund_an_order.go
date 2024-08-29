package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) RefundAnOrder(
	id string,
	payload RefundAnOrderPayload,
) (RefundAnOrderResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return RefundAnOrderResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest([]string{"1.0", "orders", id, "refund"}, http.MethodPost, body, url.Values{})
	if err != nil {
		return RefundAnOrderResponse{}, err
	}

	var response RefundAnOrderResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return RefundAnOrderResponse{}, err
	}

	return response, nil
}
