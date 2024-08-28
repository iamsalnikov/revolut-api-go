# revolut-api-go

[![Go Reference](https://pkg.go.dev/badge/github.com/jerethom/revolut-api-go.svg)](https://pkg.go.dev/github.com/jerethom/revolut-api-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/jerethom/revolut-api-go)](https://goreportcard.com/report/github.com/jerethom/revolut-api-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`revolut-api-go` is a Go package for easy interaction with the Revolut API. It provides a simple interface to perform common operations on the Revolut API, such as customer management, order creation, and webhook management.

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
)

// Default client (production)
client := revolut.NewClient("your-api-key")

// Sandbox mode client
client := revolut.NewClient("your-api-key", revolut.WithIsSandbox(true))
```

### Creating a Customer

```go
import (
    "fmt"
    "github.com/jerethom/revolut-api-go"
    "github.com/jerethom/revolut-api-go/types/customer_types"
)

func main() {
    client := revolut.NewClient("your-api-key", revolut.WithIsSandbox(true))
    
    payload := customer_types.CreateCustomerPayload{
        FullName: "John Doe",
        Email:    "john.doe@example.com",
        // Add other required fields
    }
    
    response, err := client.CreateCustomer(payload)
    if err != nil {
        fmt.Printf("Error creating customer: %v\n", err)
        return
    }
    
    fmt.Printf("Customer created successfully. ID: %s\n", response.ID)
}
```

### Creating an Order

```go
import (
    "fmt"
    "github.com/jerethom/revolut-api-go"
    "github.com/jerethom/revolut-api-go/types/order_types"
)

func main() {
    client := revolut.NewClient("your-api-key", revolut.WithIsSandbox(true))
    
    order := order_types.CreateOrderPayload{
        Amount:   100,
        Currency: "EUR",
        // Add other required fields
    }
    
    response, err := client.CreateOrder(order)
    if err != nil {
        fmt.Printf("Error creating order: %v\n", err)
        return
    }
    
    fmt.Printf("Order created successfully. ID: %s\n", response.ID)
}
```

### Creating a Webhook

```go
import (
    "fmt"
    "github.com/jerethom/revolut-api-go"
    "github.com/jerethom/revolut-api-go/types/webhook_types"
)

func main() {
    client := revolut.NewClient("your-api-key", revolut.WithIsSandbox(true))
    
    payload := webhook_types.CreateWebhookPayload{
        URL:    "https://your-webhook-url.com",
        Events: []string{"ORDER_COMPLETED", "REFUND_EXECUTED"},
    }
    
    response, err := client.CreateWebhook(payload)
    if err != nil {
        fmt.Printf("Error creating webhook: %v\n", err)
        return
    }
    
    fmt.Printf("Webhook created successfully. ID: %s\n", response.ID)
}
```

## Client Options

- `WithApiVersion(version RevolutApiVersion)`: Sets the API version to use. The latest version is used by default.
- `WithIsSandbox(sandbox bool)`: Enables or disables sandbox mode. Production mode is used by default.

## Features

- Customer management
- Order creation and management
- Webhook configuration
- And more...

## Documentation

For complete Revolut API documentation, please refer to the [official Revolut API documentation](https://developer.revolut.com/docs/api-reference/).

For documentation specific to this package, use the `go doc` command or check the [online documentation](https://pkg.go.dev/github.com/jerethom/revolut-api-go).

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.