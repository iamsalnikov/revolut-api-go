package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) CreateAWebhook(
	payload CreateAWebhookPayload,
) (CreateAWebhookResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return CreateAWebhookResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest(
		[]string{"webhooks"},
		http.MethodPost,
		body,
		url.Values{},
	)
	if err != nil {
		return CreateAWebhookResponse{}, err
	}

	response := CreateAWebhookResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return CreateAWebhookResponse{}, err
	}

	return response, nil
}
