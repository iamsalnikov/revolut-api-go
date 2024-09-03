package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) RetrieveAWebhookList() (ListWebhooksResponse, error) {
	res, err := c.clientRequest.MakeRequest(
		[]string{"1.0", "webhooks"},
		http.MethodGet,
		make([]byte, 0), url.Values{},
	)
	if err != nil {
		return nil, err
	}

	var response ListWebhooksResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
