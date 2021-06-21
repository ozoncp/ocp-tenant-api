package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	createCounter *prometheus.CounterVec
	updateCounter *prometheus.CounterVec
	removeCounter *prometheus.CounterVec
)

func RegisterMetrics() {
	createCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "create_count", // metric name
			Help: "Number of successful created tenants.",
		},
		[]string{"operation"}, // labels
	)

	updateCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "update_count", // metric name
			Help: "Number of successful created tenants.",
		},
		[]string{"operation"}, // labels
	)

	removeCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "remove_count", // metric name
			Help: "Number of successful removed tenants.",
		},
		[]string{"operation"}, // labels
	)
}

func CreateCounterInc(operation string) {
	if createCounter != nil {
		createCounter.With(prometheus.Labels{"operation": operation}).Inc()
	}
}

func UpdateCounterInc(operation string) {
	if updateCounter != nil {
		updateCounter.With(prometheus.Labels{"operation": operation}).Inc()
	}
}

func RemoveCounterInc(operation string) {
	if removeCounter != nil {
		removeCounter.With(prometheus.Labels{"operation": operation}).Inc()
	}
}
