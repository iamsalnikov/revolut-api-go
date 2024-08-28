package main

const (
	defaultMerchantApiUrl        = "https://merchant.revolut.com/api"
	defaultMerchantSandboxApiUrl = "https://sandbox-merchant.revolut.com/api"
)

type RevolutApiVersion string

const (
	RevolutApiVersionLatest   RevolutApiVersion = "2024-09-01"
	RevolutApiVersion20240901 RevolutApiVersion = "2024-09-01"
)

func (r RevolutApiVersion) String() string {
	return string(r)
}

type RevolutHeader string

const (
	RevolutHeaderApiVersion RevolutHeader = "Revolut-Api-Version"
)

func (r RevolutHeader) String() string {
	return string(r)
}

type HttpHeader string

const (
	HttpHeaderAuthorization HttpHeader = "Authorization"
	HttpHeaderContentType   HttpHeader = "Content-Type"
	HttpHeaderAccept        HttpHeader = "Accept"
)

func (h HttpHeader) String() string {
	return string(h)
}

type HttpHeaderValue string

const (
	HttpHeaderValueTypeJson   HttpHeaderValue = "application/json; charset=utf-8"
	HttpHeaderValueTypeBearer HttpHeaderValue = "Bearer %s"
)

func (h HttpHeaderValue) String() string {
	return string(h)
}
