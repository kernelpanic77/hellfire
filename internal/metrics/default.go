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

func InitDefaultMetrics(registry *MetricsRegistry, sinks map[MetricType]Sink) {
	registry.NewMetrics(Rate, "checks", Time, sinks[Rate], "default")
	registry.NewMetrics(Counter, "data_received", Data, sinks[Counter], "default")
	registry.NewMetrics(Counter, "data_sent", Data, sinks[Counter], "default")
	registry.NewMetrics(Counter, "dropped_iterations", Data, sinks[Counter], "default")
	registry.NewMetrics(Trend, "iteration_duration", Time, sinks[Trend], "default")
	registry.NewMetrics(Counter, "iterations", Data, sinks[Counter], "default")
	registry.NewMetrics(Guage, "vus", Data, sinks[Guage], "default")
	registry.NewMetrics(Guage, "vus_max", Data, sinks[Guage], "default")
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

func InitHTTPMetrics(registry *MetricsRegistry, sinks map[MetricType]Sink) {
	registry.NewMetrics(Trend, "http_req_blocked", Data, sinks[Trend], "http")
	registry.NewMetrics(Trend, "http_req_connecting", Data, sinks[Trend], "http")
	registry.NewMetrics(Trend, "http_req_duration", Data, sinks[Trend], "http")
	registry.NewMetrics(Rate, "http_req_failed", Data, sinks[Rate], "http")
	registry.NewMetrics(Trend, "http_req_receiving", Time, sinks[Trend], "http")
	registry.NewMetrics(Trend, "http_req_sending", Data, sinks[Trend], "http")
	registry.NewMetrics(Trend, "http_req_tls_handshaking", Data, sinks[Trend], "http")
	registry.NewMetrics(Trend, "http_req_waiting", Data, sinks[Trend], "http")
	registry.NewMetrics(Counter, "http_reqs", Data, sinks[Counter], "http")
}