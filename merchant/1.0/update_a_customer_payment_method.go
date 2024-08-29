package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) UpdateACustomerPaymentMethod(
	customerId, paymentMethodId string,
	payload UpdateACustomerPaymentMethodPayload,
) (UpdateACustomerPaymentMethodResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return UpdateACustomerPaymentMethodResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest(
		[]string{"1.0", "customers", customerId, "payment-methods", paymentMethodId},
		http.MethodPatch,
		body,
		url.Values{},
	)
	if err != nil {
		return UpdateACustomerPaymentMethodResponse{}, err
	}

	var response UpdateACustomerPaymentMethodResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return UpdateACustomerPaymentMethodResponse{}, err
	}

	return response, nil
}
