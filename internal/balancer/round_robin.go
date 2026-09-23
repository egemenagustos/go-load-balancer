package balancer

import "sync"

type RoundRobin struct {
	current int
	mu      sync.Mutex
}

func NewRoundRobin() *RoundRobin {
	return &RoundRobin{}
}

func (r *RoundRobin) Next(servers []*Server) *Server {
	r.mu.Lock()

	defer r.mu.Unlock()

	if len(servers) == 0 {
		return nil
	}

	for range servers {
		index := r.current % len(servers)

		server := servers[index]

		r.current++

		if server.IsHealthy() {
			return server
		}
	}

	return nil
}
