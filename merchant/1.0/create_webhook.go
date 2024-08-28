package revolut_merchant

import (
	"encoding/json"
	"github.com/jerethom/revolut-api-go/merchant/1.0/types/webhook_types"
	"net/http"
)

func (c *Merchant) CreateWebhook(
	payload webhook_types.CreateWebhookPayload,
) (webhook_types.CreateWebhookResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return webhook_types.CreateWebhookResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest([]string{"webhooks"}, http.MethodPost, body)
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
