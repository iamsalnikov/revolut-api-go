package main

import "net/http"

func (c *Client) DeleteWebhook(id string) error {
	_, err := c.makeRequest(
		[]string{"webhooks", id},
		http.MethodDelete,
		make([]byte, 0),
	)
	if err != nil {
		return err
	}
	return nil
}
