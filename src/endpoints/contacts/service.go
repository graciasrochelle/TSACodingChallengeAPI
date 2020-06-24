package contacts

import (
	"TSACodingChallengeAPI/src/common"
	"net/http"
)

type Service interface {
	Get() (resp ContactsResponse, status int, e error)
}

type service struct {
	config common.Config
}

func NewService(config common.Config) Service {
	return &service{
		config: config,
	}
}

func (s *service) Get() (resp ContactsResponse, status int, e error) {
	return resp, http.StatusOK, nil
}
