package main

import (
	"go-load-balancer/internal/balancer"
	"log"
	"net/http"
	"net/url"
)

func main() {
	url1, err := url.Parse(
		"http://localhost:8081",
	)

	if err != nil {
		log.Fatal(err)
	}
	url2, err := url.Parse(
		"http://localhost:8082",
	)
	if err != nil {
		log.Fatal(err)
	}

	url3, err := url.Parse(
		"http://localhost:8083",
	)
	if err != nil {
		log.Fatal(err)
	}

	servers := []*balancer.Server{
		{URL: url1},
		{URL: url2},
		{URL: url3},
	}

	selector := balancer.NewRoundRobin()

	lb := balancer.NewLoadBalancer(servers, selector)

	log.Printf("load balancer listening on :8080")

	if err := http.ListenAndServe(":8080", lb); err != nil {
		log.Fatal(err)
	}

}
