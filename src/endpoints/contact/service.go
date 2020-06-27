package contact

import (
	"TSACodingChallengeAPI/src/common"
	"TSACodingChallengeAPI/src/storage"
	"TSACodingChallengeAPI/src/utils"
	"net/http"

	"github.com/gofrs/uuid"
)

type Service interface {
	Post(request ContactRequest) (resp ContactResponse, status int, e error)
}

type service struct {
	config         common.Config
	storageService storage.Service
}

func NewService(config common.Config, storageService storage.Service) Service {
	return &service{config, storageService}
}

func (s *service) Post(request ContactRequest) (resp ContactResponse, status int, e error) {
	contactID := getUUID()
	contact := common.SQLContact{
		ContactID: contactID,
		FullName:  utils.NameToTitleCase(request.FullName),
		Email:     request.Email,
	}

	e = s.storageService.CreateContact(contact)
	if e != nil {
		status = http.StatusInternalServerError
		return
	}

	for _, p := range request.PhoneNumbers {
		phoneNumber := common.SQLPhoneNumber{
			PhoneID:     getUUID(),
			ContactID:   contactID,
			PhoneNumber: utils.NormalizePhoneNumber(p),
		}
		e = s.storageService.CreatePhoneNumber(phoneNumber)
		if e != nil {
			status = http.StatusInternalServerError
			return
		}
	}

	return resp, http.StatusCreated, nil
}

func getUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
