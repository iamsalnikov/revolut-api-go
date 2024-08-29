package revolut_merchant

import (
	"net/http"
	"net/url"
)

func (c *Merchant) DeleteAWebhook(id string) error {
	_, err := c.clientRequest.MakeRequest(
		[]string{"webhooks", id},
		http.MethodDelete,
		make([]byte, 0), url.Values{},
	)
	if err != nil {
		return err
	}
	return nil
}
