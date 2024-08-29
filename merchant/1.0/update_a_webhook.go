package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) UpdateAWebhook(
	webhookId string,
	payload UpdateAWebhookPayload,
) (UpdateAWebhookResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return UpdateAWebhookResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest(
		[]string{"1.0", "webhooks", webhookId},
		http.MethodPut,
		body,
		url.Values{},
	)
	if err != nil {
		return UpdateAWebhookResponse{}, err
	}

	var response UpdateAWebhookResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return UpdateAWebhookResponse{}, err
	}

	return response, nil
}
