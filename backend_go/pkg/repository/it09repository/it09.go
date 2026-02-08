package it09repository

import (
	"context"
	"test_project/pkg/model/it09model"
	"time"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	GetComments() (*[]it09model.Comment, error)
	InsertComment(comment it09model.Comment) error
}

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(dbconn *gorm.DB) ICommentRepository {
	return &CommentRepository{
		DB: dbconn,
	}
}

func (r *CommentRepository) withTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, 30*time.Second)
}

func (r *CommentRepository) GetComments() (*[]it09model.Comment, error) {
	ctx, cancel := r.withTimeout(context.Background())
	defer cancel()

	var comments []it09model.Comment
	if err := r.DB.WithContext(ctx).Table("it09").Find(&comments).Error; err != nil {
		return nil, err
	}
	return &comments, nil
}

func (r *CommentRepository) InsertComment(comment it09model.Comment) error {
	ctx, cancel := r.withTimeout(context.Background())
	defer cancel()

	if err := r.DB.WithContext(ctx).Table("it09").Create(&comment).Error; err != nil {
		return err
	}

	return nil
}
