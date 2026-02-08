package it08repository

import (
	"context"
	"test_project/pkg/model/it08model"
	"time"

	"gorm.io/gorm"
)

type IQuestionRepository interface {
	GetQuestions() (*[]it08model.Question, error)
	InsertQuestion(question it08model.Question) error
	GetNextSortOrder() (*int, error)
	DeleteQuestion(id string) error
	Reorder() error
}

type QuestionRepository struct {
	DB *gorm.DB
}

func NewQuestionRepository(dbconn *gorm.DB) IQuestionRepository {
	return &QuestionRepository{
		DB: dbconn,
	}
}

func (r *QuestionRepository) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func (r *QuestionRepository) GetQuestions() (*[]it08model.Question, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var qs []it08model.Question
	if err := r.DB.WithContext(ctx).Table("it08").Order("sort_order ASC").Find(&qs).Error; err != nil {
		return nil, err
	}

	return &qs, nil
}

func (r *QuestionRepository) InsertQuestion(question it08model.Question) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.DB.WithContext(ctx).Table("it08").Create(&question).Error; err != nil {
		return err
	}

	return nil
}

func (r *QuestionRepository) GetNextSortOrder() (*int, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var max int
	if err := r.DB.WithContext(ctx).Table("it08").
		Select("IFNULL(MAX(sort_order),0)").Scan(&max).Error; err != nil {
		return nil, err
	}

	max += 1
	return &max, nil
}

func (r *QuestionRepository) DeleteQuestion(id string) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.DB.WithContext(ctx).
		Table("it08").
		Where("id = ?", id).
		Delete(nil).Error; err != nil {
		return err
	}

	return nil
}

func (r *QuestionRepository) Reorder() error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	if err := r.DB.WithContext(ctx).
		Exec("SET @n := 0").Error; err != nil {
		return err
	}

	if err := r.DB.WithContext(ctx).
		Exec(`
			UPDATE it08
			SET sort_order = (@n := @n + 1)
			ORDER BY sort_order ASC
		`).Error; err != nil {
		return err
	}

	return nil
}
