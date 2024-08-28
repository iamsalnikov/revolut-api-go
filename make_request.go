package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) makeRequest(
	path []string,
	method string,
	body []byte,
) ([]byte, error) {
	resUrl, err := url.JoinPath(c.url.String(), path...)
	if err != nil {
		return nil, err
	}
	payload := bytes.NewReader(body)

	client := &http.Client{}
	req, err := http.NewRequest(method, resUrl, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HttpHeaderContentType.String(), HttpHeaderValueTypeJson.String())
	req.Header.Add(HttpHeaderAccept.String(), HttpHeaderValueTypeJson.String())
	req.Header.Add(HttpHeaderAuthorization.String(), fmt.Sprintf(HttpHeaderValueTypeBearer.String(), c.privateKey))
	req.Header.Add(RevolutHeaderApiVersion.String(), c.apiVersion.String())

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
