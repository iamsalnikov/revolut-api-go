package revolut_merchant

import "github.com/jerethom/revolut-api-go/shared"

type Merchant struct {
	clientRequest *shared.ClientRequest
}

func NewMerchant(clientRequest *shared.ClientRequest) *Merchant {
	return &Merchant{clientRequest: clientRequest}
}
