package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) RetrieveAnOrder(orderId string) (RetrieveAnOrderResponse, error) {
	body, err := c.clientRequest.MakeRequest(
		[]string{"orders", orderId},
		http.MethodGet,
		make([]byte, 0),
		url.Values{},
	)
	if err != nil {
		return RetrieveAnOrderResponse{}, err
	}

	var response RetrieveAnOrderResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return RetrieveAnOrderResponse{}, err
	}

	return response, nil
}
