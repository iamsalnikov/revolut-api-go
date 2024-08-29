package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) RetrieveACustomerList() (RetrieveACustomerListResponse, error) {
	res, err := c.clientRequest.MakeRequest(
		[]string{"1.0", "customers"},
		http.MethodPost,
		make([]byte, 0),
		url.Values{},
	)
	if err != nil {
		return RetrieveACustomerListResponse{}, err
	}

	var response RetrieveACustomerListResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return RetrieveACustomerListResponse{}, err
	}

	return response, nil
}
