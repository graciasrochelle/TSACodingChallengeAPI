package contact_test

import (
	"TSACodingChallengeAPI/src/endpoints/contact"
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestBind(t *testing.T) {
	testCases := map[string]struct {
		request      contact.ContactRequest
		hasError     bool
		errorMessage string
	}{
		"successful bind": {
			request: contact.ContactRequest{
				FullName:     "Test User",
				PhoneNumbers: []string{"+61 (0)414 570776"},
			},
			hasError: false,
		},
		"empty full name": {
			hasError:     true,
			errorMessage: "fullName not specified",
		},
		"atleast 1 phonenumber": {
			request: contact.ContactRequest{
				FullName:     "Test User",
				PhoneNumbers: []string{},
			},
			hasError:     true,
			errorMessage: "should have at least one phoneNumber",
		},
	}

	for tc, tp := range testCases {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(tp.request)

		req, _ := http.NewRequest("POST", "/", b)
		reqt, e := contact.Bind(req)
		actualBindedRequest := contact.ContactRequest{}
		actualBindedRequest = reqt.(contact.ContactRequest)

		if e == nil && !reflect.DeepEqual(tp.request, actualBindedRequest) {
			t.Errorf("For test case <%s>, Expected binded params are: %v, but actual params are: %v", tc, tp.request, actualBindedRequest)
		}

		if (e == nil) == (tp.hasError) {
			t.Errorf("For test case <%s>, Expected hasError is: %v, but actual result is different", tc, tp.hasError)
		}

		if e != nil && (tp.errorMessage != e.Error()) {
			t.Errorf("For test case <%s>, Expected error message is: %s, but actual message is:%s.", tc, tp.errorMessage, e.Error())
		}
	}
}
