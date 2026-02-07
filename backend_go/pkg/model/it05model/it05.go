package it05model

import "time"

type Queue struct {
	ID           string `gorm:"primaryKey"`
	CurrentQueue string `gorm:"column:current_queue;type:varchar(10)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type QueueResponse struct {
	CurrentQueue string `json:"current_queue"`
	Message      string `json:"message"`
}

type ClearQueueResponse struct {
	Message string `json:"message"`
}
