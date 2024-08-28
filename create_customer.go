package revolut_api_go

import (
	"encoding/json"
	"net/http"
)

type CreateCustomerQuery struct {
	FullName     string `json:"full_name,omitempty"`
	BusinessName string `json:"business_name,omitempty"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
}

type CreateCustomerResponse struct {
	Id           string `json:"id"`
	FullName     string `json:"full_name"`
	BusinessName string `json:"business_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func (c *Client) CreateCustomer(
	input CreateCustomerQuery,
) (CreateCustomerResponse, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return CreateCustomerResponse{}, err
	}

	res, err := c.makeRequest([]string{"customers"}, http.MethodPost, body)
	if err != nil {
		return CreateCustomerResponse{}, err
	}

	response := new(CreateCustomerResponse)
	if err := json.Unmarshal(res, response); err != nil {
		return CreateCustomerResponse{}, err
	}

	return *response, nil
}
