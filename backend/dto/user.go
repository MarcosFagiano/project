package dto

type UserRegister struct {
	GivenName   string `json:"given_name" binding:"required"`
	FamilyName  string `json:"family_name" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
