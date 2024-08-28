package webhook_types

type CreateWebhookPayload struct {
	//Your webhook's URL to which event notifications will be sent.
	//
	//Must be a valid HTTP or HTTPS URL, capable of receiving POST requests.
	Url string `json:"url"`
	//Possible values: [ORDER_COMPLETED, ORDER_AUTHORISED, ORDER_CANCELLED, ORDER_PAYMENT_AUTHENTICATED, ORDER_PAYMENT_DECLINED, ORDER_PAYMENT_FAILED], >= 1
	//
	//List of event types that the webhook is configured to listen to.
	Events []string `json:"events"`
}

type CreateWebhookResponse struct {
	//The ID of the webhook.
	Id string `json:"id"`
	//Your webhook's URL to which event notifications will be sent.
	//
	//Must be a valid HTTP or HTTPS URL, capable of receiving POST requests.
	Url string `json:"url"`
	//Possible values: [ORDER_COMPLETED, ORDER_AUTHORISED, ORDER_CANCELLED, ORDER_PAYMENT_AUTHENTICATED, ORDER_PAYMENT_DECLINED, ORDER_PAYMENT_FAILED], >= 1
	//
	//List of event types that the webhook is configured to listen to.
	Events []string `json:"events"`
	//The signing secret for the webhook. Use it to verify the signature for the webhook request's payload.
	SigningSecret string `json:"signing_secret"`
}

type ListWebhooksResponse struct {
	//The ID of the webhook.
	Id string `json:"id"`
	//Your webhook's URL to which event notifications will be sent.
	//
	//Must be a valid HTTP or HTTPS URL, capable of receiving POST requests.
	Url string `json:"url"`
	//Possible values: [ORDER_COMPLETED, ORDER_AUTHORISED, ORDER_CANCELLED, ORDER_PAYMENT_AUTHENTICATED, ORDER_PAYMENT_DECLINED, ORDER_PAYMENT_FAILED], >= 1
	//
	//List of event types that the webhook is configured to listen to.
	Events []string `json:"events"`
}
