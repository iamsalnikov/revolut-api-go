---
layout: default
title: API Reference
---

# API Reference

This page provides an overview of the main types and methods available in the revolut-api-go package.

## Client

The main entry point for interacting with the Revolut API.

```go
type Client struct {
    Merchant *revolut_merchant.Merchant
}

func NewClient(privateKey string, options ...ClientOption) *Client
```

### Client Options

- `WithApiVersion(version RevolutApiVersion)`: Sets the API version to use.
- `WithIsSandbox(sandbox bool)`: Enables or disables sandbox mode.

## Merchant

The `Merchant` struct provides methods for interacting with the Merchant API.

### Customer Operations

- `CreateACustomer(payload CreateCustomerPayload) (CreateACustomerResponse, error)`
- `RetrieveACustomer(id string) (RetrieveACustomerResponse, error)`
- `UpdateACustomer(id string, payload UpdateACustomerPayload) (UpdateACustomerResponse, error)`
- `DeleteACustomer(id string) error`

### Order Operations

- `CreateAnOrder(payload CreateAnOrderPayload) (CreateAnOrderResponse, error)`
- `RetrieveAnOrder(id string) (RetrieveAnOrderResponse, error)`
- `UpdateAnOrder(id string, payload UpdateAnOrderPayload) (UpdateAnOrderResponse, error)`
- `CancelAnOrder(id string) (CancelAnOrderResponse, error)`
- `CaptureAnOrder(id string, payload CaptureAnOrderPayload) (CaptureAnOrderResponse, error)`
- `RefundAnOrder(id string, payload RefundAnOrderPayload) (RefundAnOrderResponse, error)`

### Webhook Operations

- `CreateAWebhook(payload CreateWebhookPayload) (CreateAWebhookResponse, error)`
- `RetrieveAWebhook(webhookId string) (RetrieveAWebhookResponse, error)`
- `UpdateAWebhook(webhookId string, payload UpdateAWebhookPayload) (UpdateAWebhookResponse, error)`
- `DeleteAWebhook(id string) error`

For detailed information on the parameters and return types of these methods, please refer to the [Go Package Documentation](https://pkg.go.dev/github.com/jerethom/revolut-api-go).

## Types

The package includes various types for request payloads and response structures. Some key types include:

- `CreateCustomerPayload`
- `CreateAnOrderPayload`
- `CreateWebhookPayload`
- `Order`
- `Customer`
- `Webhook`

For a complete list of available types and their fields, consult the Go Package Documentation
