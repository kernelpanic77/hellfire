package client

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"time"

	"github.com/kernelpanic77/hellfire/internal/metrics"
)

type HttpObject struct {
	req *http.Request
	resp *http.Response
	
	time_req_start time.Time
	// idle time to obtain connection from the pool
	time_connection_blocked int64 

	// connection establish metrics
	time_for_rcv_connection int64
	time_for_tls_handshake int64  
	time_for_dns int64  
	time_for_establish_connection int64

	// time taken to send the request 
	time_for_req_sending int64
	// also known as TTFB, time till first byte  
	time_for_req_waiting int64

	time_till_first_byte time.Time
	// time taken to receive the entire response 
	time_for_recv_response int64 

	request_status bool 
}

func (h *HttpObject) getHttpTrace() *httptrace.ClientTrace {
	// variables to track time/duration
	var (
		dnsStart, dnsEnd, connStart,
		connEnd, connectStart, connectEnd,
		tlsHandShakeStart, tlsHandShakeEnd, 
		start_waiting_time, start_write_headers time.Time
	)

	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			connStart = time.Now()
		},
		GotConn: func(info httptrace.GotConnInfo) {
			connEnd = time.Now()

			if info.Reused {
				//fmt.Println("connection reused")
			} else {
				h.time_for_rcv_connection = connEnd.Sub(connStart).Milliseconds()
			}
		},
		ConnectStart: func(network, addr string) {
			connectStart = time.Now()
		},
		ConnectDone: func(network string, addr string, err error) {
			connectEnd = time.Now()
			if err != nil {
				//fmt.Println("error at ConnectDone", err)

			} else {
				h.time_for_establish_connection = connectEnd.Sub(connectStart).Milliseconds()
			}
		},
		DNSStart: func(info httptrace.DNSStartInfo) {
			dnsStart = time.Now()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			dnsEnd = time.Now()
			h.time_for_dns = dnsEnd.Sub(dnsStart).Milliseconds()
		},
		TLSHandshakeStart: func() {
			tlsHandShakeStart = time.Now()
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			if err != nil {
				//fmt.Println("tls error", err)
			} else {
				tlsHandShakeEnd = time.Now()
				h.time_for_tls_handshake = tlsHandShakeEnd.Sub(tlsHandShakeStart).Milliseconds()
			}
		},
		PutIdleConn: func(err error) {
			if err != nil {
				//fmt.Println("error at putIdleConn", err)
			} else {
				//fmt.Println("put idle connection")
			}
		},
		WroteHeaders: func() {
			start_write_headers = time.Now()
		},
		WroteRequest: func(w httptrace.WroteRequestInfo) {
			if w.Err == nil {
				h.time_for_req_sending = time.Since(start_write_headers).Milliseconds() 
				start_waiting_time = time.Now()
			}
		},
		GotFirstResponseByte: func() {
			h.time_for_req_waiting = time.Since(start_waiting_time).Milliseconds()
			h.time_till_first_byte = time.Now() 
		}, 
	}
	return trace
}

type Client struct {
	url        *url.URL
	httpClient *http.Client
	logger     *log.Logger
	samples metrics.SampleContainer
}


func (h *HttpObject) GetHttpReqObjectSize() (bodySize int64, headersSize int, err error) {
	// Calculate body size
	if h.req.Body != nil {
		bodySize, _ = io.Copy(io.Discard, h.req.Body)
	} else {
		bodySize = 0
	} // Use io.Discard instead of ioutil.Discard
	for key, values := range h.req.Header {
		headersSize += len(key) + 2 // Include ": "
		for _, value := range values {
			headersSize += len(value) + 2 // Include "\r\n" for each value
		}
	}
	headersSize += 2 // Final CRLF after headers
	return
}

func (h *HttpObject) getBodyHeaderSize(typ string) (bodySize int64, headerSize int64) {
	var body io.ReadCloser
	var header http.Header 
	
	if typ == "req" {
		body = h.req.Body 
		header = h.req.Header 
	} else if (typ == "resp") {
		body = h.resp.Body
		header = h.resp.Header 
	}

	if body != nil {
		bodySize, _ = io.Copy(io.Discard, h.resp.Body)
	} else {
		bodySize = 0
	}  // Use io.Discard instead of ioutil.Discard
	// Calculate headers size
	for key, values := range header {
		headerSize += int64(len(key) + 2) // Include ": "
		for _, value := range values {
			headerSize += int64(len(value) + 2) // Include "\r\n" for each value
		}
	}
	headerSize += 2 // Final CRLF after headers
	// objectReqSize = float64(bodySize + int64(headerSize))
	return 
}

func (h *HttpObject) GetHttpObjectSize() (objectReqSize float64, objectRespSize float64, err error) {
	// Calculate Response body size
	bodyReqSize, headersReqSize := h.getBodyHeaderSize("req")
	bodyRespSize, headersRespSize := h.getBodyHeaderSize("resp")
	objectReqSize = float64(bodyReqSize + int64(headersReqSize))
	objectRespSize = float64(bodyRespSize + int64(headersRespSize))
	return 
}

func NewClient() (Client, error) {
	c := Client{
		url:        &url.URL{},
		httpClient: http.DefaultClient,
		logger:     log.Default(),
		samples: metrics.Samples{},
	}
	return c, nil
}

func (c *Client) PrepareIterationMetricSamples(h *HttpObject) {
	send_data, recv_data, _ := h.GetHttpObjectSize()
	registry := metrics.RegistryOfRegistry["test_name"] 
	timestamp := time.Now()
	send_data_sample := metrics.NewSample(registry.FetchMetricByName("send_data"), timestamp, "send_data", send_data)
	recv_data_sample := metrics.NewSample(registry.FetchMetricByName("recv_data"), timestamp, "recv_data", recv_data)
	time_for_rcv_connection := metrics.NewSample(registry.FetchMetricByName("time_for_rcv_connection"), timestamp, "time_for_rcv_connection", float64(h.time_for_rcv_connection / 1000))
	time_for_tls_handshake := metrics.NewSample(registry.FetchMetricByName("time_for_tls_handshake"), timestamp, "time_for_tls_handshake", float64(h.time_for_tls_handshake / 1000))  
	time_for_dns := metrics.NewSample(registry.FetchMetricByName("time_for_dns"), timestamp, "time_for_dns", float64(h.time_for_dns / 1000))  
	time_for_establish_connection := metrics.NewSample(registry.FetchMetricByName("time_for_establish_connection"), timestamp, "time_for_establish_connection", float64(h.time_for_establish_connection / 1000)) 
	time_for_req_sending := metrics.NewSample(registry.FetchMetricByName("time_for_req_sending"), timestamp, "time_for_req_sending", float64(h.time_for_req_sending / 1000))
	time_for_req_waiting := metrics.NewSample(registry.FetchMetricByName("time_for_req_waiting"), timestamp, "time_for_req_waiting", float64(h.time_for_req_waiting / 1000))
	time_for_recv_response := metrics.NewSample(registry.FetchMetricByName("time_for_recv_response"), timestamp, "time_for_recv_response", float64(h.time_for_recv_response / 1000)) 
	c.samples = metrics.Samples{send_data_sample, recv_data_sample, time_for_rcv_connection, time_for_tls_handshake, time_for_dns, time_for_establish_connection, time_for_req_sending, time_for_req_waiting, time_for_recv_response}
}

func (c *Client) CollectMetrics() metrics.SampleContainer {
	curr_samples := c.samples
	c.samples = metrics.Samples{}
	return curr_samples
}

func (c *Client) Request(method string, url string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, nil)
	var resp *http.Response = nil 
	h := &HttpObject{}
	trace_obj := h.getHttpTrace()
	if err != nil {
		h.request_status = false
		c.PrepareIterationMetricSamples(h)
		return nil, err
	}

	req := request.WithContext(httptrace.WithClientTrace(request.Context(), trace_obj))
	h.time_req_start = time.Now() 
	req_client := &http.Client{Transport: &http.Transport{}}
	resp, err = req_client.Do(req)
	if err != nil {
		h.request_status = false
		c.PrepareIterationMetricSamples(h)
		return nil, err
	}

	h.time_for_recv_response = int64(time.Since(h.time_till_first_byte))
	h.req = req 
	h.resp = resp
	h.request_status = true
	c.PrepareIterationMetricSamples(h)
	return resp, nil
}