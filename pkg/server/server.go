package server

import (
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server interface {
	Serve()
}

type server struct{}

func New() Server {
	return &server{}
}

func (s *server) Serve() {
	h := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			rh(ctx)
		case "/metrics":
			ph(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
	fasthttp.ListenAndServe(":8080", h)
}

func init() {
	reg.MustRegister(requestCounter)
	reg.MustRegister(requestBytesReceived)
}

var (
	reg            = prometheus.NewRegistry()
	requestCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "fasthttp",
		Subsystem: "requests",
		Name:      "count",
	})
	requestBytesReceived = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "fasthttp",
		Subsystem: "requests",
		Name:      "received_bytes",
	})
	ph = fasthttpadaptor.NewFastHTTPHandler(promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	rh = func(ctx *fasthttp.RequestCtx) {
		requestBytesReceived.Add(float64(len(ctx.Request.Header.RawHeaders()) + len(ctx.Request.Body())))
		fmt.Fprintf(ctx, "H")
		requestCounter.Inc()
	}
)
