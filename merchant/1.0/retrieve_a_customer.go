package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) RetrieveACustomer(customerId string) (RetrieveACustomerResponse, error) {
	res, err := c.clientRequest.MakeRequest(
		[]string{"1.0", "customers", customerId},
		http.MethodGet,
		make([]byte, 0),
		url.Values{},
	)
	if err != nil {
		return RetrieveACustomerResponse{}, err
	}

	var response RetrieveACustomerResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
