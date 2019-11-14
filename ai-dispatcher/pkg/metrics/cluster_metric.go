package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	clusterModelTimeGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: "alameda_ai_dispatcher",
		Name:      "cluster_model_seconds",
		Help:      "Target modeling time of cluster",
	}, []string{"name", "data_granularity"})

	clusterModelTimeCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "alameda_ai_dispatcher",
		Name:      "cluster_model_seconds_total",
		Help:      "Total target modeling time of cluster",
	}, []string{"name", "data_granularity"})

	clusterMAPEGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: "alameda_ai_dispatcher",
		Name:      "cluster_metric_mape",
		Help:      "MAPE of cluster metric",
	}, []string{"name", "metric_type", "data_granularity"})

	clusterRMSEGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: "alameda_ai_dispatcher",
		Name:      "cluster_metric_rmse",
		Help:      "RMSE of cluster metric",
	}, []string{"name", "metric_type", "data_granularity"})

	clusterMetricDriftCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "alameda_ai_dispatcher",
		Name:      "cluster_metric_drift_total",
		Help:      "Total number of cluster metric drift",
	}, []string{"name", "data_granularity"})
)

type clusterMetric struct{}

func newClusterMetric() *clusterMetric {
	return &clusterMetric{}
}

func (clusterMetric *clusterMetric) setClusterMetricModelTime(
	name, dataGranularity string, val float64) {
	clusterModelTimeGauge.WithLabelValues(name, dataGranularity).Set(val)
}

func (clusterMetric *clusterMetric) addClusterMetricModelTimeTotal(
	name, dataGranularity string, val float64) {
	clusterModelTimeCounter.WithLabelValues(name, dataGranularity).Add(val)
}

func (clusterMetric *clusterMetric) setClusterMetricMAPE(
	name, metricType, dataGranularity string, val float64) {
	clusterMAPEGauge.WithLabelValues(name,
		metricType, dataGranularity).Set(val)
}

func (clusterMetric *clusterMetric) setClusterMetricRMSE(
	name, metricType, dataGranularity string, val float64) {
	clusterRMSEGauge.WithLabelValues(name,
		metricType, dataGranularity).Set(val)
}

func (clusterMetric *clusterMetric) addClusterMetricDrift(
	name, dataGranularity string, val float64) {
	clusterMetricDriftCounter.WithLabelValues(name, dataGranularity).Add(val)
}
