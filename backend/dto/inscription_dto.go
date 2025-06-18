package dto

type InscriptionDTO struct {
	ID          uint   `json:"id"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	ActivityIDs []uint `json:"activity_ids"` // o detalles si querés expandir
}
