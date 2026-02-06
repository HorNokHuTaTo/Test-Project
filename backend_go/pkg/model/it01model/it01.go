package it01model

type Person struct {
	ID        string `gorm:"primaryKey" json:"id"`
	FullName  string `json:"full_name"`
	BirthDate string `json:"birth_date"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
}

type Req struct {
	FullName  string `json:"full_name"`
	BirthDate string `json:"birth_date"`
	Address   string `json:"address"`
}
