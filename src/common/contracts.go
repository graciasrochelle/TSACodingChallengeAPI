package common

// ErrorResponse
//
// swagger:response
type ErrorResponse struct {
	// The error message
	Message string `json:"message"`
}

type Contacts struct {
	Contacts []SQLContact
}

type PhoneNumbers struct {
	PhoneNumbers []SQLPhoneNumber
}

type SQLPhoneNumber struct {
	PhoneID     string `json:"phoneID"`
	ContactID   string `json:"contactId"`
	PhoneNumber string `json:"phoneNumber"`
}

type SQLContact struct {
	ContactID string `json:"contactId"`
	FullName  string `json:"fullName"`
	Email     string `json:"email"`
}
