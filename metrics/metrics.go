package metrics

import(
	"github.com/prometheus/client_golang/prometheus"
)

var UploadCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "image_uploaded_total",
		Help: "total uploaded images",
	},
)

var UploadDuration = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name: "upload_duration_seconds",
		Help: "Upload duration",
	},
)
func Init() {
	prometheus.MustRegister(UploadCounter)
	prometheus.MustRegister(UploadDuration)
}