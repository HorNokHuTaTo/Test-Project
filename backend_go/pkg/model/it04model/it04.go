package it04model

import "time"

type Profile struct {
	ID         string    `gorm:"type:char(36);primaryKey" json:"id"`
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Profile    string    `json:"profile"`
	Birthday   time.Time `json:"birthday"`
	Occupation string    `json:"occupation"`
	CreatedAt  time.Time `json:"created_at"`
}

type InsertProfileRequest struct {
	ID         string `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Profile    string `json:"profile"`
	Birthday   string `json:"birthday"`
	Occupation string `json:"occupation"`
}

type Occupation struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Value string `json:"value"`
}

var Occupations = []Occupation{
	{ID: "1", Label: "Software Developer", Value: "Software Developer"},
	{ID: "2", Label: "UI/UX Designer", Value: "UI/UX Designer"},
	{ID: "3", Label: "QA / Tester", Value: "QA / Tester"},
	{ID: "4", Label: "Project Manager", Value: "Project Manager"},
	{ID: "5", Label: "System Analyst", Value: "System Analyst"},
	{ID: "6", Label: "DevOps Engineer", Value: "DevOps Engineer"},
}
