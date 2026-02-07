package it05repository

import (
	"context"
	"test_project/pkg/model/it05model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IQueueRepository interface {
	GetCurrentQueue() (*it05model.Queue, error)
	UpdateCurrentQueue(queue it05model.Queue) error
	InitializeQueue() error
}

type QueueRepository struct {
	DB *gorm.DB
}

func NewQueueRepository(dbconn *gorm.DB) IQueueRepository {
	return &QueueRepository{
		DB: dbconn,
	}
}

func (r *QueueRepository) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func (r *QueueRepository) GetCurrentQueue() (*it05model.Queue, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var queue it05model.Queue
	result := r.DB.Table("it05").WithContext(ctx).First(&queue)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &queue, nil
}

func (r *QueueRepository) UpdateCurrentQueue(queue it05model.Queue) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	result := r.DB.Table("it05").WithContext(ctx).Save(&queue)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *QueueRepository) InitializeQueue() error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var count int64
	result := r.DB.Table("it05").WithContext(ctx).Count(&count)
	if result.Error != nil {
		return result.Error
	}

	if count == 0 {
		newQueue := it05model.Queue{
			ID:           uuid.New().String(),
			CurrentQueue: "A0",
		}

		result = r.DB.Table("it05").WithContext(ctx).Create(&newQueue)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
