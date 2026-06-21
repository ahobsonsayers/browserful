package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Dashboard returns a reverse proxy to the
// agent-browser dashboard on 127.0.0.1:4848.
func Dashboard() http.Handler {
	return httputil.NewSingleHostReverseProxy(
		&url.URL{
			Scheme: "http",
			Host:   "127.0.0.1:4848",
		},
	)
}
