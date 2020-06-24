package main

import (
	"TSACodingChallengeAPI/src/common"
	"TSACodingChallengeAPI/src/endpoints/contact"
	"TSACodingChallengeAPI/src/endpoints/contacts"
	"TSACodingChallengeAPI/src/endpoints/healthCheck"
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

func NewRouter(config common.Config) *mux.Router {
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

func getRoutes(config common.Config) *Routes {
	contactsService := contacts.NewService(config)
	contactService := contact.NewService(config)

	return &Routes{
		Route{
			"Health Check",
			"GET",
			"/",
			common.GetHandler(healthCheck.NewEndpoint()),
		},
		// swagger:route GET /contacts
		//
		//
		// Get list of all contacts
		//
		//     Produces:
		//     - application/json
		//
		//     Schemes: https
		//
		//     Responses:
		//       200: ContactsResponse
		//		 400: BadRequestError
		//		 500: InteralServiceError
		//		 502: BadGatewayError
		Route{
			"Get list of all contacts",
			"GET",
			"/contacts",
			common.GetHandler(contacts.NewEndpoint(contactsService)),
		},
		// swagger:route POST /contact ContactRequest
		//
		// Create or Update contact
		//
		//     Produces:
		//     - application/json
		//
		//     Schemes: https
		//
		//     Responses:
		//       201: ContactResponse
		//       400: ValidationError
		//       500: InternalServerError
		//		 502: BadGateway
		Route{
			"Create or Update contact",
			"POST",
			"/contact",
			common.GetHandler(contact.NewEndpoint(contactService)),
		},
	}
}
