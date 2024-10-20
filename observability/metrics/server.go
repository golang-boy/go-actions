package metrics

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

type ServiceMetricsBuilder struct {
	Namespace string
	Subsystem string
}

func (b *ServiceMetricsBuilder) Build() grpc.UnaryServerInterceptor {
	reqGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: b.Namespace,
		Subsystem: b.Subsystem,
		Name:      "requests",
		Help:      "当前正在处理请求",
		ConstLabels: map[string]string{
			"service": "grpc",
		},
	}, []string{"service"})

	prometheus.MustRegister(reqGauge)

	errCnt := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: b.Namespace,
		Subsystem: b.Subsystem,
		Name:      "cnt",
		Help:      "请求错误计数",
		ConstLabels: map[string]string{
			"component": "server",
		},
	}, []string{"service"})

	response := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: b.Namespace,
		Subsystem: b.Subsystem,
		Name:      "response_time",
		Help:      "请求响应时间",
		ConstLabels: map[string]string{
			"service": "grpc",
		},
	}, []string{"service"})

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		start := time.Now()
		reqGauge.WithLabelValues(info.FullMethod).Add(1)
		defer func() {
			reqGauge.WithLabelValues(info.FullMethod).Add(-1)
			if err != nil {
				errCnt.WithLabelValues(info.FullMethod).Add(1)
			}

			response.WithLabelValues(info.FullMethod).Observe(time.Since(start).Seconds())
		}()

		resp, err = handler(ctx, req)

		return
	}
}
