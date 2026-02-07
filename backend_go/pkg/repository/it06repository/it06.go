package it06repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type ISkuRepository interface {
}

type SkuRepository struct {
	DB *gorm.DB
}

func NewSkuRepository(dbconn *gorm.DB) ISkuRepository {
	return &SkuRepository{
		DB: dbconn,
	}
}

func (r *SkuRepository) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}
