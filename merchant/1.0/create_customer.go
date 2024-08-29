package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) CreateCustomer(
	input CreateCustomerPayload,
) (CreateCustomerResponse, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return CreateCustomerResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest(
		[]string{"customers"},
		http.MethodPost,
		body,
		url.Values{},
	)
	if err != nil {
		return CreateCustomerResponse{}, err
	}

	response := new(CreateCustomerResponse)
	if err := json.Unmarshal(res, response); err != nil {
		return CreateCustomerResponse{}, err
	}

	return *response, nil
}
