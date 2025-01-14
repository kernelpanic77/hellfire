package metrics

// init default metrics
// checks	Rate
// data_received	Counter
// data_sent	Counter
// dropped_iterations	Counter
// iteration_duration	Trend
// iterations	Counter
// vus	Gauge	Current
// vus_max	Gauge

func InitDefaultMetrics(registry *MetricsRegistry) {
	// registry.RegisterMetric(Rate, "checks", Time, sinks[Rate], "default")
	registry.RegisterMetric(Counter, "recv_data", Data, "default")
	registry.RegisterMetric(Counter, "send_data", Data, "default")
	registry.RegisterMetric(Counter, "dropped_iterations", Data, "default")
	// registry.RegisterMetric(Trend, "iteration_duration", Time, sinks[Trend], "default")
	registry.RegisterMetric(Counter, "iterations", Data, "default")
	// registry.RegisterMetric(Guage, "vus", Data, sinks[Guage], "default")
	// registry.RegisterMetric(Guage, "vus_max", Data, sinks[Guage], "default")
}

// init http metrics
// http_req_blocked	Trend float 
// http_req_connecting	Trend float 
// http_req_duration	Trend float 
// http_req_failed	Rate float 
// http_req_receiving	Trend float 
// http_req_sending	Trend float 
// http_req_tls_handshaking	Trend float 	
// http_req_waiting	Trend	 float
// http_reqs	Counter	

func InitHTTPMetrics(registry *MetricsRegistry) {
	// registry.RegisterMetric(Trend, "http_req_blocked", Data, sinks[Trend], "http")
	// registry.RegisterMetric(Trend, "http_req_connecting", Data, sinks[Trend], "http")
	// registry.RegisterMetric(Trend, "http_req_duration", Data, sinks[Trend], "http")
	// registry.RegisterMetric(Rate, "http_req_failed", Data, sinks[Rate], "http")
	// registry.RegisterMetric(Trend, "http_req_receiving", Time, sinks[Trend], "http")
	// registry.RegisterMetric(Trend, "http_req_sending", Data, sinks[Trend], "http")
	// registry.RegisterMetric(Trend, "http_req_tls_handshaking", Data, sinks[Trend], "http")
	// registry.RegisterMetric(Trend, "http_req_waiting", Data, sinks[Trend], "http")

	registry.RegisterMetric(Guage, "time_connection_blocked", Time, "http")
	registry.RegisterMetric(Guage, "time_for_rcv_connection", Time, "http") 
	registry.RegisterMetric(Guage, "time_for_tls_handshake", Time, "http")
	registry.RegisterMetric(Guage, "time_for_dns", Time, "http")
	registry.RegisterMetric(Guage, "time_for_establish_connection", Time, "http")
	registry.RegisterMetric(Guage, "time_for_req_sendng", Time, "http")
	registry.RegisterMetric(Guage, "time_for_req_waiting", Time, "http")
	registry.RegisterMetric(Guage, "time_for_recv_response", Time, "http")
	
	// idle time to obtain connection from the pool
	registry.RegisterMetric(Trend, "time_connection_blocked", Time, "http")
	registry.RegisterMetric(Trend, "time_for_rcv_connection", Time, "http") 
	registry.RegisterMetric(Trend, "time_for_tls_handshake", Time, "http")
	registry.RegisterMetric(Trend, "time_for_dns", Time, "http")
	registry.RegisterMetric(Trend, "time_for_establish_connection", Time, "http")
	registry.RegisterMetric(Trend, "time_for_req_sendng", Time, "http")
	registry.RegisterMetric(Trend, "time_for_req_waiting", Time, "http")
	registry.RegisterMetric(Trend, "time_for_recv_response", Time, "http")
	// time_connection_blocked int64 
	// time_for_rcv_connection int64
	// time_for_tls_handshake int64  
	// time_for_dns int64  
	// time_for_establish_connection int64
	// time_for_req_sending int64
	// time_for_req_waiting int64
	// time_for_recv_response int64 

	registry.RegisterMetric(Counter, "http_reqs", Data, "http")
}