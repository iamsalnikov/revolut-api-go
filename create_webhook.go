package revolut_api_go

import (
	"encoding/json"
	"net/http"
)

type CreateWebhookPayload struct {
	//Your webhook's URL to which event notifications will be sent.
	//
	//Must be a valid HTTP or HTTPS URL, capable of receiving POST requests.
	Url string `json:"url"`
	//Possible values: [ORDER_COMPLETED, ORDER_AUTHORISED, ORDER_CANCELLED, ORDER_PAYMENT_AUTHENTICATED, ORDER_PAYMENT_DECLINED, ORDER_PAYMENT_FAILED], >= 1
	//
	//List of event types that the webhook is configured to listen to.
	Events []string `json:"events"`
}

func (c CreateWebhookPayload) MarshalBinary() (data []byte, err error) {
	return json.Marshal(data)
}

type CreateWebhookResponse struct {
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
	//The signing secret for the webhook. Use it to verify the signature for the webhook request's payload.
	SigningSecret string `json:"signing_secret"`
}

func (c *Client) CreateWebhook(
	payload CreateWebhookPayload,
) (CreateWebhookResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return CreateWebhookResponse{}, err
	}

	res, err := c.makeRequest([]string{"webhooks"}, http.MethodPost, body)
	if err != nil {
		return CreateWebhookResponse{}, err
	}

	response := new(CreateWebhookResponse)
	err = json.Unmarshal(res, response)
	if err != nil {
		return CreateWebhookResponse{}, err
	}

	return *response, nil
}
