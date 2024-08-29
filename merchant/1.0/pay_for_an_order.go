package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) PayForAnOrder(
	orderId string,
	payload PayForAnOrderPayload,
) (PayForAnOrderResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return PayForAnOrderResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest(
		[]string{"orders", orderId, "payments"},
		http.MethodPost,
		body,
		url.Values{},
	)
	if err != nil {
		return PayForAnOrderResponse{}, err
	}

	var response PayForAnOrderResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return PayForAnOrderResponse{}, err
	}

	return response, nil
}
