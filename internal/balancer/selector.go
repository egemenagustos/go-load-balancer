package balancer

type Selector interface {
	Next(servers []*Server) *Server
}
