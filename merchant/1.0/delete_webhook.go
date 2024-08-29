package revolut_merchant

import "net/http"

func (c *Merchant) DeleteWebhook(id string) error {
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
