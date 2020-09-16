package prometheus

import "github.com/prometheus/client_golang/prometheus"

//create your metrics here
type Metric struct {
	HttpRequestsTotal *prometheus.CounterVec
}
