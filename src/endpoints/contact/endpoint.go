package contact

import (
	"TSACodingChallengeAPI/src/common"
	"TSACodingChallengeAPI/src/utils"
	"bytes"
	"encoding/json"
	"errors"
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
	common.EncodeJsonResponse(w, httpStatus, response)
	return nil
}

func validate(req ContactRequest) error {
	if len(req.FullName) == 0 {
		return errors.New("fullName not specified")
	}
	if len(req.PhoneNumbers) == 0 {
		return errors.New("should have at least one phoneNumber")
	}
	for _, p := range req.PhoneNumbers {
		if isValid := utils.IsPossibleNumber(p); !isValid {
			return errors.New("phoneNumber entered is not valid")
		}
	}
	return nil
}
