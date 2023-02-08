package client

import (
	"fmt"
	"os"
	"sync"

	"github.com/valyala/fasthttp"
)

// New takes a Config and returns a client pointer
func New(c Config) Client {
	return &client{URL: c.URL, count: 10}
}

type client struct {
	URL   string
	count int
}

// Config is a struct containing the configuration parameters of a client
type Config struct {
	URL string
}

// Client is an interface
type Client interface {
	Call()
}

func (c *client) Call() {
	url := fasthttp.AcquireURI()
	url.Parse(nil, []byte(c.URL))

	hc := &fasthttp.HostClient{
		Addr:  string(url.Host()), // The host address and port must be set explicitly
		IsTLS: string(url.Scheme()) == "https",
	}

	var wg sync.WaitGroup

	for i := 0; i < c.count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req := fasthttp.AcquireRequest()
			req.SetURI(url) // copy url into request

			req.Header.SetMethod(fasthttp.MethodGet)
			resp := fasthttp.AcquireResponse()
			err := hc.Do(req, resp)
			fasthttp.ReleaseRequest(req)
			if err == nil {
				fmt.Printf("Response: %s\n", resp.Body())
			} else {
				fmt.Fprintf(os.Stderr, "Connection error: %v\n", err)
			}
			fasthttp.ReleaseResponse(resp)
		}()
	}
	wg.Wait()
	fasthttp.ReleaseURI(url) // now you may release the URI
}
