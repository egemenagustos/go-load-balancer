package balancer

import (
	"net/http"
	"net/http/httputil"
)

type LoadBalancer struct {
	servers  []*Server
	selector Selector
}

func NewLoadBalancer(servers []*Server, selector Selector) *LoadBalancer {
	return &LoadBalancer{
		servers:  servers,
		selector: selector,
	}
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	server := lb.selector.Next(lb.servers)

	if server == nil {
		http.Error(
			w,
			"no backend available!",
			http.StatusServiceUnavailable,
		)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(server.URL)

	proxy.ServeHTTP(w, r)
}
