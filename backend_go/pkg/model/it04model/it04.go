package it04model

import (
	"mime/multipart"
	"time"
)

type Profile struct {
	ID         string    `gorm:"type:char(36);primaryKey" json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Profile    string    `json:"profile"`
	Birthday   time.Time `json:"birthday"`
	Occupation string    `json:"occupation"`
	Sex        string    `json:"sex"`
	CreatedAt  time.Time `json:"created_at"`
}

type InsertProfileRequest struct {
	ID         string                `json:"id"`
	FirstName  string                `json:"first_name"`
	LastName   string                `json:"last_name"`
	Email      string                `json:"email"`
	Phone      string                `json:"phone"`
	Profile    *multipart.FileHeader `json:"profile"`
	Birthday   string                `json:"birthday"`
	Occupation string                `json:"occupation"`
	Sex        string                `json:"sex"`
}

type Occupation struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}

const (
	SexMale   = "Male"
	SexFemale = "Female"
)

func IsValidGender(s string) bool {
	return s == SexMale || s == SexFemale
}
