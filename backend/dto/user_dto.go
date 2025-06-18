package dto

type UserDTO struct {
	ID          uint   `json:"id"`
	GivenName   string `json:"given_name"`
	FamilyName  string `json:"family_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Admin       bool   `json:"admin"`
}
