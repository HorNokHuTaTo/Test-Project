package it08model

import (
	"time"

	"gorm.io/datatypes"
)

type Question struct {
	ID           string         `json:"id" gorm:"primaryKey"`
	QuestionText string         `json:"question_text" gorm:"column:question_text"`
	Choices      datatypes.JSON `json:"choices" gorm:"column:choices"`
	SortOrder    int            `json:"sort_order" gorm:"column:sort_order"`
	CreatedAt    time.Time      `json:"created_at" gorm:"column:created_at"`
}

type InsertQuestionRequest struct {
	QuestionText string   `json:"question_text"`
	Choices      []string `json:"choices"`
}

type QuestionResponse struct {
	ID           string    `json:"id"`
	QuestionText string    `json:"question_text"`
	Choices      []string  `json:"choices"`
	CreatedAt    time.Time `json:"created_at"`
}
