package it03model

import "time"

type Document struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateStatusRequest struct {
	IDs         []string `json:"ids"`
	Status      string   `json:"status"`
	Description string   `json:"description"`
}

type GetDocumentsResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

const (
	StatusPending  = "รออนุมัติ"
	StatusApproved = "อนุมัติ"
	StatusRejected = "ไม่อนุมัติ"
)

func IsValidStatus(s string) bool {
	return s == StatusApproved || s == StatusRejected
}
