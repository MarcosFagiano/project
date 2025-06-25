package dto

type InscriptionCreate struct {
	UserID     uint `json:"user_id" binding:"required"`
	ActivityID uint `json:"activity_id" binding:"required"`
}

type InscriptionResponse struct {
	ID         uint `json:"id"`
	UserID     uint `json:"user_id"`
	ActivityID uint `json:"activity_id"`
}
