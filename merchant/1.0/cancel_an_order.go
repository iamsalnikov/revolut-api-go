package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) CancelAnOrder(id string) (CancelAnOrderResponse, error) {
	res, err := c.clientRequest.MakeRequest(
		[]string{"orders", id, "cancel"},
		http.MethodDelete,
		make([]byte, 0),
		url.Values{},
	)
	if err != nil {
		return CancelAnOrderResponse{}, err
	}

	response := CancelAnOrderResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return CancelAnOrderResponse{}, err
	}

	return response, nil
}
