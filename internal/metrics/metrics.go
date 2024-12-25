package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "devbulls_counter",
		Help: "Counting the total number of requets being handled",
	})

	gauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "devbulls_gauge",
		Help: "Monitoring node usage",
	}, []string{"node", "namespace"})

	MetricsHandler = promhttp.Handler()
)

// func RecordMetrics() {
// 	go func() {
// 		for {
// 			counter.Inc()
// 			gauge.WithLabelValues("node-1", "namespace-b").Set(rand.Float64())
// 			time.Sleep(time.Second * 5)
// 		}
// 	}()
// }

func init() {
	prometheus.MustRegister(counter, gauge)
}
