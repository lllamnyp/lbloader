package client

import (
	"sync"
	"time"

	"github.com/valyala/fasthttp"
)

// New takes a Config and returns a client pointer
func New(c Config) Client {
	return &client{url: c.URL, count: c.Count, duration: time.Duration(c.Duration * float64(time.Second))}
}

type client struct {
	url      string
	count    int
	duration time.Duration
}

// Config is a struct containing the configuration parameters of a client
type Config struct {
	URL      string
	Count    int
	Duration float64
}

// Client is an interface
type Client interface {
	Call()
}

func (c *client) Call() {
	url := fasthttp.AcquireURI()
	url.Parse(nil, []byte(c.url))

	hc := &fasthttp.HostClient{
		Addr:  string(url.Host()), // The host address and port must be set explicitly
		IsTLS: string(url.Scheme()) == "https",
	}

	var wg sync.WaitGroup

	for worker, end := 0, time.Now().Add(c.duration); worker < c.count; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req := fasthttp.AcquireRequest()
			req.SetURI(url) // copy url into request

			req.Header.SetMethod(fasthttp.MethodGet)
			resp := fasthttp.AcquireResponse()
			for i := 0; i&0x0f != 0 || !time.Now().After(end); i++ {
				hc.Do(req, resp)
			}
			fasthttp.ReleaseRequest(req)
			fasthttp.ReleaseResponse(resp)
		}()
	}
	wg.Wait()
	fasthttp.ReleaseURI(url) // now you may release the URI
}
