package trace

import (
	"context"
	"net"
	"strings"

	"go.opentelemetry.io/otel"

	"github.com/pengcainiao2/zero/core/logx"
	oteltrace "go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc/peer"
)

const localhost = "127.0.0.1"

var (
	tracerProviders = map[string]oteltrace.Tracer{}
)

// PeerFromCtx returns the peer from ctx.
func PeerFromCtx(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	if !ok || p == nil {
		return ""
	}

	return p.Addr.String()
}

// SpanInfo returns the span info.
func SpanInfo(fullMethod, peerAddress string) (string, []attribute.KeyValue) {
	attrs := []attribute.KeyValue{RPCSystemGRPC}
	name, mAttrs := ParseFullMethod(fullMethod)
	attrs = append(attrs, mAttrs...)
	attrs = append(attrs, PeerAttr(peerAddress)...)
	return name, attrs
}

// ParseFullMethod returns the method name and attributes.
func ParseFullMethod(fullMethod string) (string, []attribute.KeyValue) {
	name := strings.TrimLeft(fullMethod, "/")
	parts := strings.SplitN(name, "/", 2)
	if len(parts) != 2 {
		// Invalid format, does not follow `/package.service/method`.
		return name, []attribute.KeyValue(nil)
	}

	var attrs []attribute.KeyValue
	if service := parts[0]; service != "" {
		attrs = append(attrs, semconv.RPCServiceKey.String(service))
	}
	if method := parts[1]; method != "" {
		attrs = append(attrs, semconv.RPCMethodKey.String(method))
	}

	return name, attrs
}

// PeerAttr returns the peer attributes.
func PeerAttr(addr string) []attribute.KeyValue {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil
	}

	if len(host) == 0 {
		host = localhost
	}

	return []attribute.KeyValue{
		semconv.NetPeerIPKey.String(host),
		semconv.NetPeerPortKey.String(port),
	}
}

func TracingEnabled(ctx context.Context, operate string) {
	if !oteltrace.SpanContextFromContext(ctx).HasTraceID() {
		logx.NewTraceLogger(ctx).Warn().Str("operate", operate).Msg("操作未包含tracing，请补充")
	}
}

func GetTracerProvider(tracerName string) oteltrace.Tracer {
	if v, ok := tracerProviders[tracerName]; ok {
		return v
	}
	tracer := otel.GetTracerProvider().Tracer(tracerName)
	tracerProviders[tracerName] = tracer
	return tracer
}
