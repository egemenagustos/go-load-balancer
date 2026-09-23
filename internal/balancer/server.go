package balancer

import (
	"net/url"
	"sync"
)

type Server struct {
	URL *url.URL

	mu        sync.RWMutex
	isHealthy bool
}

func (s *Server) IsHealthy() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.isHealthy
}

func (s *Server) SetHealthy(healthy bool) {
	s.mu.Lock()

	defer s.mu.Unlock()

	s.isHealthy = healthy
}
