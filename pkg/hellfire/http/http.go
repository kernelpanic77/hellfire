package http

import (
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	url        *url.URL
	httpClient *http.Client
	logger     *log.Logger
}

func NewClient(base string) (*Client, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	c := &Client{
		url:        baseURL,
		httpClient: http.DefaultClient,
		logger:     log.Default(),
	}
	return c, nil
}

func (c *Client) Request(method string) (*http.Response, error) {
	req := &http.Request{
		Method: method,
		URL:    c.url,
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
