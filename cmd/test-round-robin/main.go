package main

import (
	"fmt"
	"go-load-balancer/internal/balancer"
	"net/url"
)

func main() {
	url1, _ := url.Parse("http://localhost:8081")
	url2, _ := url.Parse("http://localhost:8082")
	url3, _ := url.Parse("http://localhost:8083")

	servers := []*balancer.Server{
		{URL: url1},
		{URL: url2},
		{URL: url3},
	}

	rr := balancer.NewRoundRobin()

	for i := 0; i < 10; i++ {
		server := rr.Next(servers)

		fmt.Println(server.URL)
	}
}
