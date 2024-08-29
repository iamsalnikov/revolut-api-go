package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) CreateACustomer(
	input CreateACustomerPayload,
) (CreateACustomerResponse, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return CreateACustomerResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest(
		[]string{"customers"},
		http.MethodPost,
		body,
		url.Values{},
	)
	if err != nil {
		return CreateACustomerResponse{}, err
	}

	response := CreateACustomerResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return CreateACustomerResponse{}, err
	}

	return response, nil
}
