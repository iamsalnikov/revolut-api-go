package revolut_merchant

import (
	"encoding/json"
	"net/http"
)

func (c *Merchant) CreateWebhook(
	payload CreateWebhookPayload,
) (CreateWebhookResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return CreateWebhookResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest([]string{"webhooks"}, http.MethodPost, body, url.Values{})
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
