package it09model

import "time"

type Comment struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"column:username"`
	Message   string    `json:"message" gorm:"column:message"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

type InsertCommentRequest struct {
	Message string `json:"message"`
}
