package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) RotateAWebhookSigningSecret(
	webhookId string,
	payload RotateAWebhookSigningSecretPayload,
) (RotateAWebhookSigningSecretResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return RotateAWebhookSigningSecretResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest(
		[]string{"1.0", "webhooks", webhookId, "rotate-signing-secret"},
		http.MethodPost,
		body,
		url.Values{},
	)
	if err != nil {
		return RotateAWebhookSigningSecretResponse{}, err
	}

	var response RotateAWebhookSigningSecretResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return RotateAWebhookSigningSecretResponse{}, err
	}

	return response, nil
}
