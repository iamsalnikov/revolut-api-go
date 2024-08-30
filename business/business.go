package revolut_business

import "github.com/jerethom/revolut-api-go/shared"

// Business todo
type Business struct {
	clientRequest *shared.ClientRequest
}

func NewBusiness(clientRequest *shared.ClientRequest) *Business {
	return &Business{clientRequest: clientRequest}
}
