package contacts

import (
	"TSACodingChallengeAPI/src/common"
	"TSACodingChallengeAPI/src/storage"
	"errors"
	"net/http"
)

var (
	displayContactDetails = []string{"Full Name", "Email", "Phone Numbers"}
)

type Service interface {
	Get() (resp ContactsResponse, status int, e error)
}

type service struct {
	config         common.Config
	storageService storage.Service
}

func NewService(config common.Config, storageService storage.Service) Service {
	return &service{config, storageService}
}

func (s *service) Get() (response ContactsResponse, status int, e error) {
	if response, e = s.getAllContacts(); e != nil {
		status = http.StatusInternalServerError
		return
	}
	response.DisplayContactDetails = displayContactDetails

	if response.Contacts == nil {
		status = http.StatusNoContent
		e = errors.New("No contacts available")
		return
	}

	return response, http.StatusOK, nil
}

func (s *service) getAllContacts() (response ContactsResponse, err error) {
	contacts, err := s.storageService.ReadContacts()
	if err != nil {
		return
	}

	for _, c := range contacts.Contacts {
		contact := Contact{
			ID:       c.ContactID,
			Email:    c.Email,
			FullName: c.FullName,
		}

		if phoneNums, e := s.storageService.ReadPhoneNumbers(c.ContactID); e == nil {
			for _, p := range phoneNums.PhoneNumbers {
				contact.PhoneNumbers = append(contact.PhoneNumbers, p.PhoneNumber)
			}
		}

		response.Contacts = append(response.Contacts, contact)
	}
	return
}
