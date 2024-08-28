package revolut_api_go

import (
	"encoding/json"
	"net/http"
	"revolut-api-go/types/customer_types"
)

func (c *Client) CreateCustomer(
	input customer_types.CreateCustomerPayload,
) (customer_types.CreateCustomerResponse, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return customer_types.CreateCustomerResponse{}, err
	}

	res, err := c.makeRequest([]string{"customers"}, http.MethodPost, body)
	if err != nil {
		return customer_types.CreateCustomerResponse{}, err
	}

	response := new(customer_types.CreateCustomerResponse)
	if err := json.Unmarshal(res, response); err != nil {
		return customer_types.CreateCustomerResponse{}, err
	}

	return *response, nil
}
