package common

import "net/http"

type Client interface {
	Request(method string, url string) (*http.Response, error)
}