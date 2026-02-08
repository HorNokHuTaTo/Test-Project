package it06model

import "time"

type Sku struct {
	ID        string    `json:"id"`
	SkuCode   string    `json:"sku_code"`
	Barcode   string    `json:"barcode"`
	CreatedAt time.Time `json:"created_at"`
}

type InsertSkuRequest struct {
	SkuCode string `json:"sku_code"`
}
