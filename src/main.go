package main

import (
	"TSACodingChallengeAPI/src/shared"
	"net/http"
	"os"
)

func main() {

	port := "10010"

	if os.Getenv("HTTP_PLATFORM_PORT") != "" {
		port = os.Getenv("HTTP_PLATFORM_PORT")
	}

	// Init config
	config := shared.NewConfig()

	c := initRouter(*config)
	println("Starting service at " + port)
	panic(http.ListenAndServe(":"+port, c))
}

func initRouter(config shared.Config) *Chain {
	router := NewRouter(config)

	c := NewChain()
	c.Use(router)
	c.Build()

	return c
}
