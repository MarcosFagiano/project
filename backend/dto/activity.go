package dto

import "time"

type Activity struct {
	ID          uint      `json:"id"`
	CategoryID  uint      `json:"category_id"`
	Description string    `json:"description"`
	Quota       uint      `json:"quota"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
