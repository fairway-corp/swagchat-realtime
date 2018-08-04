package apmhttprouter

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/elastic/apm-agent-go"
	"github.com/elastic/apm-agent-go/module/apmhttp"
)

// Wrap wraps h such that it will report requests as transactions
// to Elastic APM, using route in the transaction name.
//
// By default, the returned Handle will use elasticapm.DefaultTracer.
// Use WithTracer to specify an alternative tracer.
//
// By default, the returned Handle will recover panics, reporting
// them to the configured tracer. To override this behaviour, use
// WithRecovery.
func Wrap(h httprouter.Handle, route string, o ...Option) httprouter.Handle {
	opts := options{
		tracer: elasticapm.DefaultTracer,
	}
	for _, o := range o {
		o(&opts)
	}
	if opts.recovery == nil {
		opts.recovery = apmhttp.NewTraceRecovery(opts.tracer)
	}
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		if !opts.tracer.Active() {
			h(w, req, p)
			return
		}
		tx := opts.tracer.StartTransaction(req.Method+" "+route, "request")
		ctx := elasticapm.ContextWithTransaction(req.Context(), tx)
		req = apmhttp.RequestWithContext(ctx, req)
		defer tx.End()

		finished := false
		body := opts.tracer.CaptureHTTPRequestBody(req)
		w, resp := apmhttp.WrapResponseWriter(w)
		defer func() {
			if v := recover(); v != nil {
				opts.recovery(w, req, body, tx, v)
				finished = true
			}
			apmhttp.SetTransactionContext(tx, req, resp, body, finished)
		}()
		h(w, req, p)
		finished = true
	}
}

type options struct {
	tracer   *elasticapm.Tracer
	recovery apmhttp.RecoveryFunc
}

// Option sets options for tracing.
type Option func(*options)

// WithTracer returns an Option which sets t as the tracer
// to use for tracing server requests.
func WithTracer(t *elasticapm.Tracer) Option {
	if t == nil {
		panic("t == nil")
	}
	return func(o *options) {
		o.tracer = t
	}
}

// WithRecovery returns an Option which sets r as the recovery
// function to use for tracing server requests.
func WithRecovery(r apmhttp.RecoveryFunc) Option {
	if r == nil {
		panic("r == nil")
	}
	return func(o *options) {
		o.recovery = r
	}
}
