package balancer

import (
	"log"
	"net/http"
	"time"
)

type HealthChecker struct {
	servers  []*Server
	client   *http.Client
	interval time.Duration
}

func NewHealthChecker(servers []*Server, interval time.Duration) *HealthChecker {
	return &HealthChecker{
		servers: servers,
		client: &http.Client{
			Timeout: 2 * time.Second,
		},
		interval: interval,
	}
}

func (h *HealthChecker) Start() {
	h.checkAll()

	ticker := time.NewTicker(h.interval)
	defer ticker.Stop()

	for range ticker.C {
		h.checkAll()
	}
}

func (h *HealthChecker) checkAll() {
	for _, server := range h.servers {
		go h.check(server)
	}
}

func (h *HealthChecker) check(server *Server) {
	healthURL := server.URL.String() + "/health"

	resp, err := h.client.Get(healthURL)

	if err != nil {
		server.SetHealthy(false)

		log.Printf(
			"backend unhealthy: %s",
			server.URL,
		)

		return
	}

	defer resp.Body.Close()

	healthy := resp.StatusCode == http.StatusOK

	server.SetHealthy(healthy)

	log.Printf(
		"backend %s healthy=%t",
		server.URL,
		healthy,
	)
}
