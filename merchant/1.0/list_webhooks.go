package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) ListWebhooks() ([]ListWebhooksResponse, error) {
	res, err := c.clientRequest.MakeRequest(
		[]string{"webhooks"},
		http.MethodGet,
		make([]byte, 0), url.Values{},
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
