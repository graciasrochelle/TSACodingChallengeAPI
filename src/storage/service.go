package storage

import "TSACodingChallengeAPI/src/common"

// Service interface
type Service interface {
	ReadContacts() (common.Contacts, error)
	ReadPhoneNumbers(contactID string) (common.PhoneNumbers, error)
	CreateContact(contact common.SQLContact) error
	CreatePhoneNumber(phoneNumber common.SQLPhoneNumber) error
}

type service struct {
	config          common.Config
	inMemoryService InMemoryService
	sqlService      SQLService
}

// NewService creates new service
func NewService(config common.Config, inMemoryService InMemoryService, sqlService SQLService) Service {
	return &service{
		config,
		inMemoryService,
		sqlService,
	}
}

func (s *service) ReadContacts() (common.Contacts, error) {
	if s.config.UseInMemoryStorage {
		return s.inMemoryService.ReadContacts()
	}
	return s.sqlService.ReadContacts()
}

func (s *service) ReadPhoneNumbers(contactID string) (phoneNumbers common.PhoneNumbers, e error) {
	if s.config.UseInMemoryStorage {
		return s.inMemoryService.ReadPhoneNumbers(contactID)
	}
	return s.sqlService.ReadPhoneNumbers(contactID)
}

func (s *service) CreateContact(contact common.SQLContact) (e error) {
	if s.config.UseInMemoryStorage {
		return s.inMemoryService.CreateContact(contact)
	}
	return s.sqlService.CreateContact(contact)
}

func (s *service) CreatePhoneNumber(phoneNumber common.SQLPhoneNumber) (e error) {
	if s.config.UseInMemoryStorage {
		return s.inMemoryService.CreatePhoneNumber(phoneNumber)
	}
	return s.sqlService.CreatePhoneNumber(phoneNumber)
}
