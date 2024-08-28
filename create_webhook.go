package revolut

import (
	"encoding/json"
	"github.com/jerethom/revolut-api-go/types/webhook_types"
	"net/http"
)

func (c *Client) CreateWebhook(
	payload webhook_types.CreateWebhookPayload,
) (webhook_types.CreateWebhookResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return webhook_types.CreateWebhookResponse{}, err
	}

	res, err := c.makeRequest([]string{"webhooks"}, http.MethodPost, body)
	if err != nil {
		return webhook_types.CreateWebhookResponse{}, err
	}

	response := new(webhook_types.CreateWebhookResponse)
	err = json.Unmarshal(res, response)
	if err != nil {
		return webhook_types.CreateWebhookResponse{}, err
	}

	return *response, nil
}
