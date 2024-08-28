package main

import (
	"net/url"
)

type Client struct {
	privateKey string
	apiVersion RevolutApiVersion
	url        *url.URL
	isSandbox  bool
}

type ClientOption func(client *Client)

func WithApiVersion(version RevolutApiVersion) ClientOption {
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
		client.apiVersion = RevolutApiVersionLatest
	}

	if client.isSandbox {
		client.url, _ = url.Parse(defaultMerchantSandboxApiUrl)
	} else {
		client.url, _ = url.Parse(defaultMerchantApiUrl)
	}

	newPath, _ := url.JoinPath(client.url.Path, string(client.apiVersion))
	client.url.Path = newPath

	return client
}
