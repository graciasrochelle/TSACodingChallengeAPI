package main

import (
	"TSACodingChallengeAPI/src/common"
	"TSACodingChallengeAPI/src/endpoints/healthCheck"
	"TSACodingChallengeAPI/src/shared"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Routes []Route

func NewRouter(config shared.Config) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	routes := getRoutes(config)

	for _, route := range *routes {
		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			HandlerFunc(route.Handler)
	}
	return router
}

func getRoutes(config shared.Config) *Routes {
	return &Routes{
		Route{
			"Health Check",
			"GET",
			"/",
			common.GetHandler(healthCheck.NewEndpoint()),
		},
	}
}
