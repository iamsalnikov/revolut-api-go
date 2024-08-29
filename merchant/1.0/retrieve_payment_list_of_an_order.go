package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) RetrievePaymentListOfAnOrder(
	orderId string,
) (RetrievePaymentListOfAnOrderResponse, error) {
	res, err := c.clientRequest.MakeRequest(
		[]string{"orders", orderId, "payments"},
		http.MethodGet,
		make([]byte, 0),
		url.Values{},
	)
	if err != nil {
		return RetrievePaymentListOfAnOrderResponse{}, err
	}

	var response RetrievePaymentListOfAnOrderResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
