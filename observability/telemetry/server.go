package telemetry

import (
	"context"

	"go.opencensus.io/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type OtelBuilder struct {
	Tracer trace.Tracer
}

func (b *OtelBuilder) Build() grpc.UnaryServerInterceptor {

	if b.Tracer == nil {
		b.Tracer = otel.GetTracerProvider().Tracer("grpc.server")
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx = b.extract(ctx)
		spanCtx, span := b.Tracer.Start(ctx, info.FullMethod)

		span.SetAttributes(attribute.String("address", info.FullMethod))
		defer func() {
			if err != nil {
				span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
				span.RecordError(err)
			}
		}()

		resp, err = handler(spanCtx, req)
		return

	}

}

func (b *OtelBuilder) extract(ctx context.Context) context.Context {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}

	return otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(md))
}
