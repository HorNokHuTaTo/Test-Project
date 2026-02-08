package it06repository

import (
	"context"
	"test_project/pkg/model/it06model"
	"time"

	"gorm.io/gorm"
)

type ISkuRepository interface {
	GetSkus() (*[]it06model.Sku, error)
	InsertSku(sku it06model.Sku) error
	DeleteSku(id string) error
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

func (r *SkuRepository) GetSkus() (*[]it06model.Sku, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var skus []it06model.Sku

	if err := r.DB.WithContext(ctx).Table("it06").
		Scan(&skus).
		Error; err != nil {
		return nil, err
	}

	return &skus, nil
}

func (r *SkuRepository) InsertSku(sku it06model.Sku) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.DB.WithContext(ctx).Table("it06").Create(&sku).Error; err != nil {
		return err
	}

	return nil
}

func (r *SkuRepository) DeleteSku(id string) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var sku it06model.Sku
	if err := r.DB.WithContext(ctx).Table("it06").
		Where("id = ?", id).
		Delete(&sku).Error; err != nil {
		return err
	}

	return nil
}
