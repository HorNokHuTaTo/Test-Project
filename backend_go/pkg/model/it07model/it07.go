package it07model

import "time"

type Sku struct {
	ID        string    `json:"id"`
	SkuCode   string    `json:"sku_code"`
	QrCode    string    `json:"qr_code"`
	CreatedAt time.Time `json:"created_at"`
}

type InsertSkuRequest struct {
	SkuCode string `json:"sku_code"`
}
