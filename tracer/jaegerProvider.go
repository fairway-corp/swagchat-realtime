package tracer

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/swagchat/rtm-api/config"
	"github.com/swagchat/rtm-api/logger"
	jaeger "github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/transport/zipkin"
)

var (
	jaegerTracer opentracing.Tracer
	jaegerCloser io.Closer
)

type jaegerProvider struct {
	ctx context.Context

	// zipkin
	endpoint  string
	batchSize int
	timeout   int
}

func (jp *jaegerProvider) NewTracer() error {
	var tracer opentracing.Tracer
	var closer io.Closer
	var err error

	if jp.endpoint == "" {
		// jaeger
		cfg := &jaegerConfig.Configuration{
			Sampler: &jaegerConfig.SamplerConfig{
				Type:  "const",
				Param: 1,
			},
			Reporter: &jaegerConfig.ReporterConfig{
				LogSpans: true,
			},
		}
		tracer, closer, err = cfg.New(
			fmt.Sprintf("%s:%s", config.AppName, config.BuildVersion),
			jaegerConfig.Logger(jaeger.StdLogger),
		)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	} else {
		// zipkin
		transport, err := zipkin.NewHTTPTransport(
			jp.endpoint,
			zipkin.HTTPBatchSize(jp.batchSize),
			zipkin.HTTPTimeout(time.Second*time.Duration(jp.timeout)),
			zipkin.HTTPLogger(jaeger.StdLogger),
		)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		tracer, closer = jaeger.NewTracer(
			fmt.Sprintf("%s:%s", config.AppName, config.BuildVersion),
			jaeger.NewConstSampler(true),
			jaeger.NewRemoteReporter(transport),
		)
		if err != nil {
			logger.Error(err.Error())
			return err
		}
	}

	opentracing.SetGlobalTracer(tracer)

	jaegerTracer = tracer
	jaegerCloser = closer
	return nil
}

func (jp *jaegerProvider) StartTransaction(name, transactionType string, opts ...StartTransactionOption) (context.Context, interface{}) {
	if jaegerTracer == nil {
		return jp.ctx, nil
	}

	opt := startTransactionOptions{}
	for _, o := range opts {
		o(&opt)
	}

	var span opentracing.Span
	if opt.r == nil {
		span = jaegerTracer.StartSpan(name)
	} else {
		spanCtx, _ := jaegerTracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(opt.r.Header))
		span = jaegerTracer.StartSpan(name, ext.RPCServerOption(spanCtx))
	}

	ctx := opentracing.ContextWithSpan(jp.ctx, span)
	ctx = context.WithValue(ctx, config.CtxTracerSpan, span)
	return ctx, span
}

func (jp *jaegerProvider) StartSpan(name, spanType string) interface{} {
	if jaegerTracer == nil {
		return nil
	}

	span, _ := opentracing.StartSpanFromContext(jp.ctx, fmt.Sprintf("%s.%s", spanType, name))
	return span
}

func (jp *jaegerProvider) InjectHTTPRequest(span interface{}, req *http.Request) {
	if span == nil {
		return
	}

	s := span.(opentracing.Span)
	ext.SpanKindRPCClient.Set(s)
	ext.HTTPUrl.Set(s, req.RequestURI)
	ext.HTTPMethod.Set(s, req.Method)
	s.Tracer().Inject(
		s.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
}

func (jp *jaegerProvider) SetTag(span interface{}, key string, value interface{}) {
	if span == nil {
		return
	}
	span.(opentracing.Span).SetTag(key, value)
}

func (jp *jaegerProvider) SetHTTPStatusCode(span interface{}, statusCode int) {
	if span == nil {
		return
	}
	span.(opentracing.Span).SetTag("http.status_code", statusCode)
}

func (jp *jaegerProvider) SetError(span interface{}, err error) {
	if span == nil {
		return
	}
	span.(opentracing.Span).SetTag("error", true)
	span.(opentracing.Span).SetTag("message", err.Error())
}

func (jp *jaegerProvider) Finish(span interface{}) {
	if span == nil {
		return
	}
	span.(opentracing.Span).Finish()
}

func (jp *jaegerProvider) CloseTransaction() {
	span := jp.ctx.Value(config.CtxTracerSpan)
	if span == nil {
		return
	}
	span.(opentracing.Span).Finish()
}

func (jp *jaegerProvider) Close() {
	if jaegerCloser == nil {
		return
	}
	jaegerCloser.Close()
}
