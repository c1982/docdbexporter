package exporter

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusExporter struct {
	prefix   string
	registry *prometheus.Registry
}

func NewPrometheusExporter(prefix string) *PrometheusExporter {
	registry := prometheus.NewRegistry()
	e := &PrometheusExporter{
		prefix:   prefix,
		registry: registry,
	}
	registerMetrics(e)
	return e
}

func (p *PrometheusExporter) registerGaugeWithLabels(name string, help string, labels []string) *GaugeMetric {
	var a GaugeMetric
	gauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: p.prefix,
		Name:      name,
		Help:      help,
	}, labels)
	p.registry.MustRegister(gauge)
	a.gauge = *gauge

	return &a
}

func (p *PrometheusExporter) CollectMetricsPeriodically(interval time.Duration, fn func()) {
	time.AfterFunc(interval, func() {
		fn()
		p.CollectMetricsPeriodically(interval, fn)
	})
}

func (p *PrometheusExporter) ListenAndServe(addr string) error {
	http.Handle(
		"/metrics", promhttp.HandlerFor(
			p.registry,
			promhttp.HandlerOpts{
				EnableOpenMetrics: false,
			}),
	)
	return http.ListenAndServe(addr, nil)
}
