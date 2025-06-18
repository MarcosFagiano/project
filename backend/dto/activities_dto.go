package dto

type ActivityDTO struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Quota       int    `json:"quota"`
	Category    string `json:"category"`
}
