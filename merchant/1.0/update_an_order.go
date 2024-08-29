package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) UpdateAnOrder(
	orderId string,
	payload UpdateAnOrderPayload,
) (UpdateAnOrderResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return UpdateAnOrderResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest(
		[]string{"orders", orderId},
		http.MethodPatch,
		body,
		url.Values{},
	)
	if err != nil {
		return UpdateAnOrderResponse{}, err
	}

	response := UpdateAnOrderResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
