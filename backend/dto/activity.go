package dto

import "time"

type Activity struct {
	ID          uint      `json:"id"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Quota       int       `json:"quota"`
	CategoryID  uint      `json:"category_id"`
}
