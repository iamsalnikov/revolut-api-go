package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type RetrieveAllPaymentMethodOfACustomerFilter struct {
	OnlyMerchant bool
}

func (c *Merchant) RetrieveAllPaymentMethodOfACustomer(
	customerId string,
	filter RetrieveAllPaymentMethodOfACustomerFilter,
) (RetrieveAllPaymentMethodOfACustomerResponse, error) {
	params := url.Values{}
	if filter.OnlyMerchant {
		params.Add("only_merchant", "true")
	}

	res, err := c.clientRequest.MakeRequest(
		[]string{"1.0", "customers", customerId, "payment-methods"},
		http.MethodGet,
		make([]byte, 0),
		params,
	)
	if err != nil {
		return RetrieveAllPaymentMethodOfACustomerResponse{}, err
	}

	var response RetrieveAllPaymentMethodOfACustomerResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return RetrieveAllPaymentMethodOfACustomerResponse{}, err
	}

	return response, nil
}
