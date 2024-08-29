package revolut_merchant

import (
	"net/http"
	"net/url"
)

func (c *Merchant) DeleteAWebhook(webhookId string) error {
	_, err := c.clientRequest.MakeRequest(
		[]string{"webhooks", webhookId},
		http.MethodDelete,
		make([]byte, 0), url.Values{},
	)
	if err != nil {
		return err
	}
	return nil
}
