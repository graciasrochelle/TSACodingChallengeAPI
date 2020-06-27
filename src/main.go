package main

import (
	"TSACodingChallengeAPI/src/common"
	"TSACodingChallengeAPI/src/storage"
	"net/http"
	"os"
)

func main() {

	port := "10010"

	if os.Getenv("HTTP_PLATFORM_PORT") != "" {
		port = os.Getenv("HTTP_PLATFORM_PORT")
	}

	// Init config
	config := common.NewConfig()

	c := initRouter(*config)
	println("Starting service at " + port)
	panic(http.ListenAndServe(":"+port, c))
}

func initRouter(config common.Config) *Chain {
	storageService := storage.NewService(config)
	storageService.CreateConnectionPool()

	router := NewRouter(config, storageService)

	c := NewChain()
	c.Use(router)
	c.Build()

	return c
}
