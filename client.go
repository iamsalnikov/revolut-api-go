package revolut

import (
	"github.com/jerethom/revolut-api-go/constants"
	"github.com/jerethom/revolut-api-go/merchant/1.0"
	"github.com/jerethom/revolut-api-go/shared"
	"net/url"
)

type Client struct {
	privateKey string
	apiVersion constants.RevolutApiVersion
	url        *url.URL
	isSandbox  bool

	Merchant *revolut_merchant.Merchant
}

type ClientOption func(client *Client)

func WithApiVersion(version constants.RevolutApiVersion) ClientOption {
	return func(client *Client) {
		client.apiVersion = version
	}
}

func WithIsSandbox(sandbox bool) ClientOption {
	return func(client *Client) {
		client.isSandbox = sandbox
	}
}

func NewClient(privateKey string, options ...ClientOption) *Client {
	client := &Client{privateKey: privateKey}

	for _, option := range options {
		option(client)
	}

	if client.apiVersion == "" {
		client.apiVersion = constants.RevolutApiVersionLatest
	}

	if client.isSandbox {
		client.url, _ = url.Parse(constants.DefaultMerchantSandboxApiUrl)
	} else {
		client.url, _ = url.Parse(constants.DefaultMerchantApiUrl)
	}

	newPath, _ := url.JoinPath(client.url.Path, string(client.apiVersion))
	client.url.Path = newPath

	clientRequest := shared.NewClientRequest(client.privateKey, client.apiVersion, client.url)
	client.Merchant = revolut_merchant.NewMerchant(clientRequest)

	return client
}
