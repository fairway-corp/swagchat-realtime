package tracer

import (
	"context"
	"net/http"

	"github.com/swagchat/rtm-api/config"
)

type provider interface {
	NewTracer() error
	StartTransaction(name, transactionType string, opts ...StartTransactionOption) (context.Context, interface{})
	StartSpan(name, spanType string) interface{}
	SetTag(span interface{}, key string, value interface{})
	SetHTTPStatusCode(span interface{}, statusCode int)
	SetError(span interface{}, err error)
	Finish(span interface{})
	CloseTransaction()
	Close()
}

type startTransactionOptions struct {
	r *http.Request
}

type StartTransactionOption func(*startTransactionOptions)

func StartTransactionOptionWithHTTPRequest(r *http.Request) StartTransactionOption {
	return func(ops *startTransactionOptions) {
		ops.r = r
	}
}

func Provider(ctx context.Context) provider {
	cfg := config.Config()
	var p provider

	switch cfg.Tracer.Provider {
	case "jaeger":
		p = &jaegerProvider{
			ctx: ctx,
		}
	case "elasticapm":
		p = &elasticapmProvider{
			ctx: ctx,
		}
	default:
		p = &notuseProvider{
			ctx: ctx,
		}
	}

	return p
}
