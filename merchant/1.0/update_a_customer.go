package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) UpdateACustomer(
	customerId string,
	payload UpdateACustomerPayload,
) (UpdateACustomerResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return UpdateACustomerResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest(
		[]string{"1.0", "customers", customerId},
		http.MethodPatch,
		body,
		url.Values{},
	)
	if err != nil {
		return UpdateACustomerResponse{}, err
	}

	var response UpdateACustomerResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
