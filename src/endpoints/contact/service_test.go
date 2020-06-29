package contact_test

import (
	"TSACodingChallengeAPI/src/common"
	"TSACodingChallengeAPI/src/endpoints/contact"
	"TSACodingChallengeAPI/src/storage"
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestAddContactService(t *testing.T) {
	config := common.Config{
		UseInMemoryStorage: true,
		Contacts: []common.SQLContact{
			common.SQLContact{
				ContactID: "39358d6f-88a0-42d4-8b8c-26719d7a9b5f",
				FullName:  "Radia Perlman",
				Email:     "rper1001@mit.edu",
			},
		},
		PhoneNumbers: []common.SQLPhoneNumber{
			common.SQLPhoneNumber{
				PhoneID:     "39458d6f-88a0-65d4-8b8c-26719d7h9b5f",
				ContactID:   "39358d6f-88a0-42d4-8b8c-26719d7a9b5f",
				PhoneNumber: "+61 (0)414 570776",
			},
		},
	}
	testCases := map[string]struct {
		request     contact.ContactRequest
		expResponse contact.ContactResponse
		status      int
		hasError    bool
	}{
		"successful response": {
			request: contact.ContactRequest{
				FullName:     "Test User",
				PhoneNumbers: []string{"+61 (0)414 570776"},
				Email:        "test.user@test.com",
			},
			expResponse: contact.ContactResponse{},
			status:      http.StatusCreated,
			hasError:    false,
		},
	}
	for tc, tp := range testCases {
		inMemoryService := storage.NewInMemoryService(config)
		storageService := storage.NewService(config, inMemoryService, &mockedSQLService{})
		cs := contact.NewService(config, storageService)
		resp, status, e := cs.Post(tp.request)
		if e == nil && !reflect.DeepEqual(tp.expResponse, resp) {
			t.Errorf("For test case <%s>, Expected response is: %v, but actual response was: %v", tc, tp.expResponse, resp)
		}

		if (e == nil) == (tp.hasError) {
			t.Errorf("For test case <%s>, Expected hasError is: %v, but actual result is different", tc, tp.hasError)
		}

		if tp.status != status {
			t.Errorf("For test case <%s>, Expected status code is: %s, but actual status code is:%s.", tc, tp.status, status)
		}
	}
}

type mockedSQLService struct {
}

func (s *mockedSQLService) CreateConnectionPool() (err error) {
	return errors.New("mocked sql server")
}
func (s *mockedSQLService) ReadContacts() (c common.Contacts, e error) {
	return c, e
}
func (s *mockedSQLService) ReadPhoneNumbers(contactID string) (p common.PhoneNumbers, e error) {
	return p, e
}
func (s *mockedSQLService) CreateContact(contact common.SQLContact) error {
	return errors.New("mocked sql server")
}
func (s *mockedSQLService) CreatePhoneNumber(phoneNumber common.SQLPhoneNumber) error {
	return errors.New("mocked sql server")
}
