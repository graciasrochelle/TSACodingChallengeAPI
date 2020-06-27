package contacts

type ContactsResponse struct {
	Contacts              []Contact `json:"contacts"`
	DisplayContactDetails []string  `json:"displayContactDetails"`
}

type Contact struct {
	ID           string   `json:"id"`
	FullName     string   `json:"fullName"`
	Email        string   `json:"email"`
	PhoneNumbers []string `json:"phoneNumbers"`
}
