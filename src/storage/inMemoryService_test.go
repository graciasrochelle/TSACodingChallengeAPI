package storage_test

import (
	"TSACodingChallengeAPI/src/common"
	"TSACodingChallengeAPI/src/storage"
	"reflect"
	"testing"
)

func TestInMemryStorageService(t *testing.T) {
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
		reqtAddContact     common.SQLContact
		reqtAddPhone       common.SQLPhoneNumber
		expGetContacts     common.Contacts
		expGetPhoneNumbers common.PhoneNumbers
	}{
		"successful response": {
			reqtAddContact: common.SQLContact{
				FullName:  "Test User",
				ContactID: "random-uuid",
				Email:     "test.user@test.com",
			},
			reqtAddPhone: common.SQLPhoneNumber{
				PhoneID:     "39458d6f-88a0",
				ContactID:   "random-uuid",
				PhoneNumber: "+61 (0)414 570776",
			},
			expGetContacts: common.Contacts{
				Contacts: []common.SQLContact{
					common.SQLContact{
						ContactID: "39358d6f-88a0-42d4-8b8c-26719d7a9b5f",
						FullName:  "Radia Perlman",
						Email:     "rper1001@mit.edu",
					},
					common.SQLContact{
						FullName:  "Test User",
						ContactID: "random-uuid",
						Email:     "test.user@test.com",
					},
				},
			},
			expGetPhoneNumbers: common.PhoneNumbers{
				PhoneNumbers: []common.SQLPhoneNumber{
					common.SQLPhoneNumber{
						PhoneID:     "39458d6f-88a0",
						ContactID:   "random-uuid",
						PhoneNumber: "+61 (0)414 570776",
					},
				},
			},
		},
	}
	for tc, tp := range testCases {
		ims := storage.NewInMemoryService(config)

		e := ims.CreateContact(tp.reqtAddContact)
		if e != nil {
			t.Errorf("For test case <%s>, returned erro <%s>", tc, e)
		}

		e = ims.CreatePhoneNumber(tp.reqtAddPhone)
		if e != nil {
			t.Errorf("For test case <%s>, returned erro <%s>", tc, e)
		}

		c, e := ims.ReadContacts()
		if e != nil {
			t.Errorf("For test case <%s>, returned erro <%s>", tc, e)
		}
		if e == nil && !reflect.DeepEqual(tp.expGetContacts, c) {
			t.Errorf("For test case <%s>, Expected response is: %v, but actual response was: %v", tc, tp.expGetContacts, c)
		}

		p, e := ims.ReadPhoneNumbers("random-uuid")
		if e != nil {
			t.Errorf("For test case <%s>, returned erro <%s>", tc, e)
		}
		if e == nil && !reflect.DeepEqual(tp.expGetPhoneNumbers, p) {
			t.Errorf("For test case <%s>, Expected response is: %v, but actual response was: %v", tc, tp.expGetPhoneNumbers, p)
		}
	}
}
