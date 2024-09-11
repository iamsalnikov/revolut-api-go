---
layout: default
title: Usage
---

# Usage

This guide will walk you through the basic usage of the revolut-api-go package.

## Client Initialization

First, you need to initialize a client with your API key:

```go
import "github.com/jerethom/revolut-api-go"

// Default client (production)
client := revolut.NewClient("your-api-key")

// Sandbox mode client
client := revolut.NewClient("your-api-key", revolut.WithIsSandbox(true))
```

## Creating a Customer

Here's an example of how to create a customer:

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
    
    response, err := client.Merchant.CreateACustomer(payload)
    if err != nil {
        fmt.Printf("Error creating customer: %v\n", err)
        return
    }
    
    fmt.Printf("Customer created successfully. ID: %s\n", response.Id)
}
```

## Creating an Order

Here's how you can create an order:

```go
import (
    "fmt"
    "github.com/jerethom/revolut-api-go"
    "github.com/jerethom/revolut-api-go/revolut_merchant"
)

func main() {
    client := revolut.NewClient("your-api-key", revolut.WithIsSandbox(true))
    
    order := revolut_merchant.CreateAnOrderPayload{
        Amount:   10000, // 100.00 in minor currency units
        Currency: "EUR",
        // Add other required fields
    }
    
    response, err := client.Merchant.CreateAnOrder(order)
    if err != nil {
        fmt.Printf("Error creating order: %v\n", err)
        return
    }
    
    fmt.Printf("Order created successfully. ID: %s\n", response.Id)
}
```

## Creating a Webhook

To create a webhook:

```go
import (
    "fmt"
    "github.com/jerethom/revolut-api-go"
    "github.com/jerethom/revolut-api-go/revolut_merchant"
)

func main() {
    client := revolut.NewClient("your-api-key", revolut.WithIsSandbox(true))
    
    payload := revolut_merchant.CreateWebhookPayload{
        Url:    "https://your-webhook-url.com",
        Events: []string{"ORDER_COMPLETED", "ORDER_AUTHORISED"},
    }
    
    response, err := client.Merchant.CreateAWebhook(payload)
    if err != nil {
        fmt.Printf("Error creating webhook: %v\n", err)
        return
    }
    
    fmt.Printf("Webhook created successfully. ID: %s\n", response.Id)
}
```

For more detailed information on available methods and their usage, please refer to the [API Reference](./api-reference) section.
