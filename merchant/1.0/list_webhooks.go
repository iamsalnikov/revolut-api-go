package revolut_merchant

import (
	"encoding/json"
	"github.com/jerethom/revolut-api-go/merchant/1.0/types/webhook_types"
	"net/http"
)

func (c *Merchant) ListWebhooks() ([]webhook_types.ListWebhooksResponse, error) {
	res, err := c.clientRequest.MakeRequest(
		[]string{"webhooks"},
		http.MethodGet,
		make([]byte, 0),
	)
	if err != nil {
		return nil, err
	}

	response := make([]webhook_types.ListWebhooksResponse, 0)
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
