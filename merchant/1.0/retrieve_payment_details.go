package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) RetrievePaymentDetails(
	paymentId string,
) (RetrievePaymentDetailsResponse, error) {
	res, err := c.clientRequest.MakeRequest(
		[]string{"payments", paymentId},
		http.MethodGet,
		make([]byte, 0),
		url.Values{},
	)
	if err != nil {
		return RetrievePaymentDetailsResponse{}, err
	}

	var response RetrievePaymentDetailsResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return RetrievePaymentDetailsResponse{}, err
	}

	return response, nil
}
