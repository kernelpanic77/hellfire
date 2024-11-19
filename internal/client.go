// interface for a generic client
// could be extended based on the usecases
// TODO: httpClient and grpcClient

package internal

import "net/http"

type Client interface {
	Request(method string, url string) (*http.Response, error)
}