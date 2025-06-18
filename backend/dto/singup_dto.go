package dto

type SingupDTO struct {
	GivenName   string `json:"given_name" binding:"required"`
	FamilyName  string `json:"family_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
