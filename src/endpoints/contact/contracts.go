package contact

type ContactRequest struct {
	ID           string   `json:"id"`
	FullName     string   `json:"fullName"`
	Email        string   `json:"email"`
	PhoneNumbers []string `json:"phoneNumbers"`
}

type ContactResponse struct {
}
