package revolut_merchant

import (
	"net/http"
	"net/url"
)

func (c *Merchant) DeleteACustomer(id string) error {
	_, err := c.clientRequest.MakeRequest(
		[]string{"1.0", "customers", id},
		http.MethodDelete,
		make([]byte, 0),
		url.Values{},
	)
	return err
}
