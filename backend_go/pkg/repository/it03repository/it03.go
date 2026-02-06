package it03repository

import (
	"context"
	"fmt"
	"test_project/pkg/model/it03model"
	"time"

	"gorm.io/gorm"
)

type IDocumentRepository interface {
	UpdateDocumentsStatus(req it03model.UpdateStatusRequest) error
	GetAllDocuments() (*[]it03model.Document, error)
}

type DocumentRepository struct {
	DB *gorm.DB
}

func NewDocumentRepository(dbconn *gorm.DB) IDocumentRepository {
	return &DocumentRepository{
		DB: dbconn,
	}
}

func (r *DocumentRepository) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func (r *DocumentRepository) UpdateDocumentsStatus(req it03model.UpdateStatusRequest) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	tx := r.DB.Table("it03").
		WithContext(ctx).
		Model(&it03model.Document{}).
		Where("id IN ?", req.IDs).
		Where("status = ?", it03model.StatusPending).
		Updates(map[string]interface{}{
			"status":      req.Status,
			"description": req.Description,
			"updated_at":  time.Now(),
		})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return fmt.Errorf("ไม่สามารถอัปเดตรายการที่อนุมัติแล้วได้")
	}

	return nil
}

func (r *DocumentRepository) GetAllDocuments() (*[]it03model.Document, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()
	var documents []it03model.Document
	err := r.DB.Table("it03").WithContext(ctx).Find(&documents).Error
	if err != nil {
		return nil, err
	}

	return &documents, nil
}
