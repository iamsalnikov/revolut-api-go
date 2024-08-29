package revolut

import "fmt"

func main() {
	client := NewClient("")

	fmt.Print(client.Merchant)
}
