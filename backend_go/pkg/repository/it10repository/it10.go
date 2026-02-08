package it10repository

import (
	"context"
	"test_project/pkg/model/it10model"
	"time"

	"gorm.io/gorm"
)

type IExamRepository interface {
	SubmitExam(result it10model.ExamResult) error
}

type ExamRepository struct {
	DB *gorm.DB
}

func NewExamRepository(dbconn *gorm.DB) IExamRepository {
	return &ExamRepository{
		DB: dbconn,
	}
}

func (r *ExamRepository) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func (r *ExamRepository) SubmitExam(result it10model.ExamResult) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.DB.WithContext(ctx).Table("it10").Create(&result).Error; err != nil {
		return err
	}

	return nil
}
