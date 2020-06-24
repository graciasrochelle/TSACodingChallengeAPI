package common

import (
	"net/http"
)

type Parameters interface{}
type ResponseType interface{}
type Handler func(r *http.Request, params Parameters) (response ResponseType, statusCode int, err error)
type Binder func(r *http.Request) (params Parameters, err error)
type ResponseEncoder func(w http.ResponseWriter, httpStatus int, response ResponseType) error
type ErrorLogger func(r *http.Request, httpStatus int, e error)

type Endpoint struct {
	handler Handler
	binder  Binder
	encoder ResponseEncoder
	logger  ErrorLogger
}

func NewEndpoint(handler Handler, binder Binder, encoder ResponseEncoder, logger ErrorLogger) *Endpoint {
	return &Endpoint{handler, binder, encoder, logger}
}

func GetHandler(api *Endpoint) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var params Parameters
		var e error
		var resp ResponseType
		var statusCode int

		if api.binder != nil {
			params, e = api.binder(r)
			if e != nil {
				statusCode = http.StatusBadRequest
			}
		}
		if e == nil {
			resp, statusCode, e = api.handler(r, params)
		}

		if e != nil {
			if statusCode < http.StatusBadRequest {
				statusCode = http.StatusInternalServerError
			}
			api.logger(r, statusCode, e)
			newErrorResponse := ErrorResponse{e.Error()}
			api.encoder(w, statusCode, newErrorResponse)
		} else {
			api.encoder(w, statusCode, resp)
		}
	})
}
