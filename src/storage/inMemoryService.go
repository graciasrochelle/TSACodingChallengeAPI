package storage

import "TSACodingChallengeAPI/src/common"

var (
	imContacts     common.Contacts
	imPhoneNumbers common.PhoneNumbers
)

// InMemoryService in-memory service interface
type InMemoryService interface {
	ReadContacts() (common.Contacts, error)
	ReadPhoneNumbers(contactID string) (common.PhoneNumbers, error)
	CreateContact(contact common.SQLContact) error
	CreatePhoneNumber(phoneNumber common.SQLPhoneNumber) error
}

type inMemoryService struct {
	config common.Config
}

// NewInMemoryService creates new in-memory service
func NewInMemoryService(config common.Config) InMemoryService {
	imContacts.Contacts = config.Contacts
	imPhoneNumbers.PhoneNumbers = config.PhoneNumbers
	return &inMemoryService{
		config: config,
	}
}

func (s *inMemoryService) ReadContacts() (common.Contacts, error) {
	return imContacts, nil
}

func (s *inMemoryService) ReadPhoneNumbers(contactID string) (phoneNumbers common.PhoneNumbers, e error) {
	phoneNumbers.PhoneNumbers = getPhoneNumbers(contactID, imPhoneNumbers.PhoneNumbers)
	return
}

func (s *inMemoryService) CreateContact(contact common.SQLContact) (e error) {
	imContacts.Contacts = append(imContacts.Contacts, contact)
	return
}

func (s *inMemoryService) CreatePhoneNumber(phoneNumber common.SQLPhoneNumber) (e error) {
	imPhoneNumbers.PhoneNumbers = append(imPhoneNumbers.PhoneNumbers, phoneNumber)
	return
}

func getPhoneNumbers(id string, phoneNumbers []common.SQLPhoneNumber) (result []common.SQLPhoneNumber) {
	for _, phoneNumber := range phoneNumbers {
		if phoneNumber.ContactID == id {
			result = append(result, phoneNumber)
		}
	}
	return result
}
