package it07repository

import (
	"context"
	"test_project/pkg/model/it07model"
	"time"

	"gorm.io/gorm"
)

type ISku2Repository interface {
	GetSkus() (*[]it07model.Sku, error)
	InsertSku(sku it07model.Sku) error
	DeleteSku(id string) error
}

type Sku2Repository struct {
	DB *gorm.DB
}

func NewSku2Repository(dbconn *gorm.DB) ISku2Repository {
	return &Sku2Repository{
		DB: dbconn,
	}
}

func (r *Sku2Repository) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func (r *Sku2Repository) GetSkus() (*[]it07model.Sku, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var skus []it07model.Sku

	if err := r.DB.WithContext(ctx).Table("it07").
		Scan(&skus).
		Error; err != nil {
		return nil, err
	}

	return &skus, nil
}

func (r *Sku2Repository) InsertSku(sku it07model.Sku) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.DB.WithContext(ctx).Table("it07").Create(&sku).Error; err != nil {
		return err
	}

	return nil
}

func (r *Sku2Repository) DeleteSku(id string) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var sku it07model.Sku
	if err := r.DB.WithContext(ctx).Table("it07").
		Where("id = ?", id).
		Delete(&sku).Error; err != nil {
		return err
	}

	return nil
}
