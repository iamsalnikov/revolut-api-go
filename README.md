# revolut-api-go

[![Go Reference](https://pkg.go.dev/badge/github.com/jerethom/revolut-api-go.svg)](https://pkg.go.dev/github.com/jerethom/revolut-api-go)
[![Go Version](https://img.shields.io/github/go-mod/go-version/jerethom/revolut-api-go)](https://github.com/jerethom/revolut-api-go/blob/main/go.mod)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/jerethom/revolut-api-go)](https://github.com/jerethom/revolut-api-go/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/jerethom/revolut-api-go)](https://goreportcard.com/report/github.com/jerethom/revolut-api-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`revolut-api-go` is a Go package that provides a client for interacting with the Revolut Merchant API. It offers a comprehensive set of functions to manage customers, orders, payments, and webhooks.

## Features

- Full support for Revolut Merchant API operations
- Sandbox mode for testing
- Customizable API version
- Type-safe request and response structures

## Installation

To install the package, use the `go get` command:

```bash
go get github.com/jerethom/revolut-api-go
```

## Usage

### Client Initialization

```go
import (
    "github.com/jerethom/revolut-api-go"
    "github.com/jerethom/revolut-api-go/constants"
)

// Default client (production)
client := revolut.NewClient("your-api-key")

// Sandbox mode client
client := revolut.NewClient("your-api-key", revolut.WithIsSandbox(true))

// Custom API version
client := revolut.NewClient("your-api-key", revolut.WithApiVersion(constants.RevolutApiVersion20240901))
```

### Creating an Order

```go
import (
    "fmt"
    "github.com/jerethom/revolut-api-go"
    "github.com/jerethom/revolut-api-go/merchant/1.0"
)

func main() {
    client := revolut.NewClient("your-api-key", revolut.WithIsSandbox(true))

    order := revolut_merchant.CreateAnOrderPayload{
        Amount:      1000, // 10.00 in minor currency units
        Currency:    "EUR",
        Description: "Test order",
    }

    response, err := client.Merchant.CreateAnOrder(order)
    if err != nil {
        fmt.Printf("Error creating order: %v\n", err)
        return
    }

    fmt.Printf("Order created successfully. ID: %s\n", response.Id)
}
```

### Managing Customers

```go
customer := revolut_merchant.CreateCustomerPayload{
    FullName: "John Doe",
    Email:    "john.doe@example.com",
    Phone:    "+1234567890",
}

response, err := client.Merchant.CreateACustomer(customer)
if err != nil {
    // Handle error
}

fmt.Printf("Customer created with ID: %s\n", response.Id)
```

### Working with Webhooks

```go
webhook := revolut_merchant.CreateWebhookPayload{
    Url:    "https://your-webhook-url.com",
    Events: []string{"ORDER_COMPLETED", "ORDER_CANCELLED"},
}

response, err := client.Merchant.CreateAWebhook(webhook)
if err != nil {
    // Handle error
}

fmt.Printf("Webhook created with ID: %s\n", response.Id)
```

## API Coverage

The package covers all major operations of the Revolut Merchant API, including:

- Customer management (create, retrieve, update, delete)
- Order operations (create, retrieve, update, cancel, capture, refund)
- Payment handling
- Webhook management

For a complete list of available methods, please refer to the [Go Reference documentation](https://pkg.go.dev/github.com/jerethom/revolut-api-go).

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.