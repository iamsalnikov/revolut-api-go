package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) RetrieveACustomerPaymentMethod(
	customerId, paymentMethodId string,
) (RetrieveACustomerPaymentMethodResponse, error) {
	res, err := c.clientRequest.MakeRequest(
		[]string{"1.0", "customers", customerId, "payment-methods", paymentMethodId},
		http.MethodGet,
		make([]byte, 0),
		url.Values{},
	)
	if err != nil {
		return RetrieveACustomerPaymentMethodResponse{}, err
	}

	var response RetrieveACustomerPaymentMethodResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return RetrieveACustomerPaymentMethodResponse{}, err
	}

	return response, nil
}
