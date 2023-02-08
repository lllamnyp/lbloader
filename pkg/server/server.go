package server

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type Server interface {
	Serve()
}

type server struct{}

func New() Server {
	return &server{}
}

func (s *server) Serve() {
	i := 0
	rh := func(ctx *fasthttp.RequestCtx) {
		i++
		fmt.Fprintf(ctx, "Hello, world! %d\n", i)
	}
	fasthttp.ListenAndServe(":8080", rh)
}
