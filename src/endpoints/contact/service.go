package contact

import (
	"TSACodingChallengeAPI/src/common"
	"net/http"
)

type Service interface {
	Post(request ContactRequest) (resp ContactResponse, status int, e error)
}

type service struct {
	config common.Config
}

func NewService(config common.Config) Service {
	return &service{
		config: config,
	}
}

func (s *service) Post(request ContactRequest) (resp ContactResponse, status int, e error) {
	return resp, http.StatusCreated, nil
}
