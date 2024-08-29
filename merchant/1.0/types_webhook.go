package revolut_merchant

type CreateAWebhookPayload struct {
	//Your webhook's URL to which event notifications will be sent.
	//
	//Must be a valid HTTP or HTTPS URL, capable of receiving POST requests.
	Url string `json:"url"`
	//Possible values: [ORDER_COMPLETED, ORDER_AUTHORISED, ORDER_CANCELLED, ORDER_PAYMENT_AUTHENTICATED, ORDER_PAYMENT_DECLINED, ORDER_PAYMENT_FAILED], >= 1
	//
	//List of event types that the webhook is configured to listen to.
	Events []string `json:"events"`
}

type Webhook struct {
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

type WebhookWithSigningSecret struct {
	Webhook
	//The signing secret for the webhook. Use it to verify the signature for the webhook request's payload.
	SigningSecret string `json:"signing_secret"`
}

type CreateAWebhookResponse struct {
	WebhookWithSigningSecret
}

type ListWebhooksResponse []Webhook

type RetrieveAWebhookResponse struct {
	WebhookWithSigningSecret
}

type UpdateAWebhookPayload struct {
	//Possible values: <= 2000 characters
	//
	//Your webhook's URL to which event notifications will be sent.
	//
	//Must be a valid HTTP or HTTPS URL, capable of receiving POST requests.
	Url string `json:"url"`
	//Possible values: [ORDER_COMPLETED, ORDER_AUTHORISED, ORDER_CANCELLED, ORDER_PAYMENT_AUTHENTICATED, ORDER_PAYMENT_DECLINED, ORDER_PAYMENT_FAILED], >= 1
	//
	//List of event types that the webhook is configured to listen to.
	Events []string `json:"events"`
}

type UpdateAWebhookResponse struct {
	Webhook
}

type RotateAWebhookSigningSecretPayload struct {
	ExpirationPeriod string `json:"expiration_period"`
}

type RotateAWebhookSigningSecretResponse struct {
	WebhookWithSigningSecret
}
