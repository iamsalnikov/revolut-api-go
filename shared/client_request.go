package shared

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/jerethom/revolut-api-go/constants"
	"io"
	"net/http"
	"net/url"
)

func NewClientRequest(privateKey string, apiVersion constants.RevolutApiVersion, url *url.URL) *ClientRequest {
	return &ClientRequest{privateKey: privateKey, apiVersion: apiVersion, url: url}
}

type ClientRequest struct {
	privateKey string
	apiVersion constants.RevolutApiVersion
	url        *url.URL
}

func (c *ClientRequest) MakeRequest(
	path []string,
	method string,
	body []byte,
	queryParams url.Values,
) ([]byte, error) {
	resUrl, err := url.JoinPath(c.url.String(), path...)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(resUrl)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for k, v := range queryParams {
		for _, vv := range v {
			q.Add(k, vv)
		}
	}
	u.RawQuery = q.Encode()

	payload := bytes.NewReader(body)

	client := &http.Client{}
	req, err := http.NewRequest(method, u.String(), payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add(constants.HttpHeaderContentType.String(), constants.HttpHeaderValueTypeJson.String())
	req.Header.Add(constants.HttpHeaderAccept.String(), constants.HttpHeaderValueTypeJson.String())
	req.Header.Add(constants.HttpHeaderAuthorization.String(), fmt.Sprintf(constants.HttpHeaderValueTypeBearer.String(), c.privateKey))
	req.Header.Add(constants.RevolutHeaderApiVersion.String(), c.apiVersion.String())

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < http.StatusOK ||
		res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New(string(body))
	}

	return body, nil
}
