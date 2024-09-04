package http

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"time"
)

func getHttpTrace() *httptrace.ClientTrace {

	// variables to track time/duration
	var (
		dnsStart, dnsEnd, connStart,
		connEnd, connectStart, connectEnd,
		tlsHandShakeStart, tlsHandShakeEnd time.Time
	)

	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			connStart = time.Now()
		},
		GotConn: func(info httptrace.GotConnInfo) {
			connEnd = time.Now()

			if info.Reused {
				log.Println("connection reused")
			} else {
				log.Println("time elapsed for Getting connection in micro seconds ", connEnd.Sub(connStart).Microseconds())

			}

		},
		ConnectStart: func(network, addr string) {
			connectStart = time.Now()

		},
		ConnectDone: func(network, addr string, err error) {
			connectEnd = time.Now()
			if err != nil {
				log.Println("error at ConnectDone", err)

			} else {
				log.Println("time elapsed to  connect  in micro seconds ", connectEnd.Sub(connectStart).Microseconds())
			}
		},
		DNSStart: func(info httptrace.DNSStartInfo) {
			dnsStart = time.Now()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			dnsEnd = time.Now()
			log.Println("time elapsed to resolve DNS in micro seconds ", dnsEnd.Sub(dnsStart).Microseconds())

		},
		TLSHandshakeStart: func() {
			tlsHandShakeStart = time.Now()
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			if err != nil {
				log.Println("tls error", err)

			} else {
				tlsHandShakeEnd = time.Now()
				log.Println("time elapsed for TLS Handshake in micro seconds ", tlsHandShakeEnd.Sub(tlsHandShakeStart).Microseconds())

			}

		},
		PutIdleConn: func(err error) {
			if err != nil {
				log.Println("error at putIdleConn", err)
			} else {
				log.Println("put idle connection")
			}

		},
	}

	return trace

}

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

func (c *Client) Request(method string, URL base) (*http.Response, error) {
	request := &http.Request{
		Method: method,
		URL:    c.url,
	}
	trace_obj := getHttpTrace()
	req := request.WithContext(httptrace.WithClientTrace(request.Context(), trace_obj))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
