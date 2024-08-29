package revolut_merchant

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Merchant) CaptureAnOrder(
	id string,
	payload CaptureAnOrderPayload,
) (CaptureAnOrderResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return CaptureAnOrderResponse{}, err
	}

	res, err := c.clientRequest.MakeRequest(
		[]string{"orders", id, "capture"},
		http.MethodPost,
		body,
		url.Values{},
	)
	if err != nil {
		return CaptureAnOrderResponse{}, err
	}

	var response CaptureAnOrderResponse
	if err := json.Unmarshal(res, &response); err != nil {
		return CaptureAnOrderResponse{}, err
	}

	return response, nil
}
