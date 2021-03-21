package main

import (
	"log"
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

func NewGlobalTracer() (opentracing.Tracer, func()) {
	cfg := jaegercfg.Configuration{
		ServiceName: "demo",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, _ := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)

	return tracer, func() { closer.Close() }
}

func main() {
	tracer, cleanup := NewGlobalTracer()

	opentracing.SetGlobalTracer(tracer)
	defer cleanup()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		clientSpan := tracer.StartSpan("demo.hello")
		defer clientSpan.Finish()

		url := "http://127.0.0.1:8080/world"
		req, _ := http.NewRequest("GET", url, nil)

		ext.SpanKindRPCClient.Set(clientSpan)
		ext.HTTPUrl.Set(clientSpan, url)
		ext.HTTPMethod.Set(clientSpan, "GET")

		tracer.Inject(clientSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
		http.DefaultClient.Do(req)
	})

	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		serverSpan := tracer.StartSpan("demo.world", ext.RPCServerOption(spanCtx))
		defer serverSpan.Finish()
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
