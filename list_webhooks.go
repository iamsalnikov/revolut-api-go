package revolut_api_go

import (
	"encoding/json"
	"net/http"
)

type ListWebhooksResponse struct {
	//The ID of the webhook.
	Id string `json:"id"`
	//Your webhook's URL to which event notifications will be sent.
	//
	//Must be a valid HTTP or HTTPS URL, capable of receiving POST requests.
	Url string `json:"url"`
	//Possible values: [ORDER_COMPLETED, ORDER_AUTHORISED, ORDER_CANCELLED, ORDER_PAYMENT_AUTHENTICATED, ORDER_PAYMENT_DECLINED, ORDER_PAYMENT_FAILED], >= 1
	//
	//List of event types that the webhook is configured to listen to.
	Events []string `json:"events"`
}

func (c *Client) ListWebhooks() ([]ListWebhooksResponse, error) {
	res, err := c.makeRequest(
		[]string{"webhooks"},
		http.MethodGet,
		make([]byte, 0),
	)
	if err != nil {
		return nil, err
	}

	response := make([]ListWebhooksResponse, 0)
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
