package it09svc

import (
	"test_project/pkg/model/it09model"
	"test_project/pkg/repository/it09repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ICommentService interface {
	GetComments() (*[]it09model.Comment, error)
	InsertComment(req it09model.InsertCommentRequest) error
}

type CommentService struct {
	CommentRepository it09repository.ICommentRepository
}

func NewCommentService(dbcon *gorm.DB) ICommentService {
	return &CommentService{
		CommentRepository: it09repository.NewCommentRepository(dbcon),
	}
}

func (s *CommentService) GetComments() (*[]it09model.Comment, error) {
	res, err := s.CommentRepository.GetComments()
	if err != nil {
		return nil, err
	}

	comments := make([]it09model.Comment, 0)
	for _, comment := range *res {
		comments = append(comments, it09model.Comment{
			ID:        comment.ID,
			Username:  comment.Username,
			Message:   comment.Message,
			CreatedAt: comment.CreatedAt,
		})
	}

	return &comments, nil
}

func (s *CommentService) InsertComment(req it09model.InsertCommentRequest) error {
	newComment := it09model.Comment{
		ID:       uuid.New().String(),
		Username: "Blend 285",
		Message:  req.Message,
	}

	if err := s.CommentRepository.InsertComment(newComment); err != nil {
		return err
	}

	return nil
}
