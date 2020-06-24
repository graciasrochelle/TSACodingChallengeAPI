package contact

import (
	"TSACodingChallengeAPI/src/common"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func NewEndpoint(s Service) *common.Endpoint {
	return common.NewEndpoint(Handle(s), Bind, Encoder)
}

func Bind(r *http.Request) (params common.Parameters, err error) {
	cr := ContactRequest{}
	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	e := json.Unmarshal(bodyBuffer, &cr)
	r.Body = ioutil.NopCloser(bytes.NewReader(bodyBuffer))
	if e != nil {
		return nil, e
	}
	return cr, validate(cr)
}

func Handle(s Service) func(r *http.Request, params common.Parameters) (response common.ResponseType, statusCode int, err error) {
	return func(r *http.Request, params common.Parameters) (response common.ResponseType, statusCode int, err error) {
		return s.Post(params.(ContactRequest))
	}
}

func Encoder(w http.ResponseWriter, httpStatus int, response common.ResponseType) error {
	w.Header().Set("Cache-Control", "no-transform, max-age=1800")
	common.EncodeJsonResponse(w, httpStatus, response)
	return nil
}

func validate(req ContactRequest) error {

	return nil
}
