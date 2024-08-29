package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) RetrieveAWebhook(webhookId string) (RetrieveAWebhookResponse, error) {
	res, err := c.clientRequest.MakeRequest(
		[]string{"1.0", "webhooks", webhookId},
		http.MethodGet,
		make([]byte, 0),
		url.Values{},
	)
	if err != nil {
		return RetrieveAWebhookResponse{}, err
	}

	var response RetrieveAWebhookResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return RetrieveAWebhookResponse{}, err
	}

	return response, nil
}
