package contacts

import (
	"TSACodingChallengeAPI/src/common"
	"net/http"
)

func NewEndpoint(s Service) *common.Endpoint {
	return common.NewEndpoint(Handle(s), Bind, Encoder)
}

func Bind(r *http.Request) (params common.Parameters, err error) {
	return
}

func Handle(s Service) func(r *http.Request, params common.Parameters) (response common.ResponseType, statusCode int, err error) {
	return func(r *http.Request, params common.Parameters) (response common.ResponseType, statusCode int, err error) {
		return s.Get()
	}
}

func Encoder(w http.ResponseWriter, httpStatus int, response common.ResponseType) error {
	w.Header().Set("Cache-Control", "no-transform, max-age=1800")
	common.EncodeJsonResponse(w, httpStatus, response)
	return nil
}
