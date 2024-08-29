package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) CancelOrder(id string) (CancelOrderResponse, error) {
	res, err := c.clientRequest.MakeRequest(
		[]string{"orders", id, "cancel"},
		http.MethodDelete,
		make([]byte, 0),
		url.Values{},
	)
	if err != nil {
		return CancelOrderResponse{}, err
	}

	response := CancelOrderResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return CancelOrderResponse{}, err
	}

	return response, nil
}
