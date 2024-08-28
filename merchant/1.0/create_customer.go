package revolut_merchant

import (
	"encoding/json"
	"github.com/jerethom/revolut-api-go/merchant/1.0/types/customer_types"
	"net/http"
)

func (c *Merchant) CreateCustomer(
	input customer_types.CreateCustomerPayload,
) (customer_types.CreateCustomerResponse, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return customer_types.CreateCustomerResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest([]string{"customers"}, http.MethodPost, body)
	if err != nil {
		return customer_types.CreateCustomerResponse{}, err
	}

	response := new(customer_types.CreateCustomerResponse)
	if err := json.Unmarshal(res, response); err != nil {
		return customer_types.CreateCustomerResponse{}, err
	}

	return *response, nil
}
