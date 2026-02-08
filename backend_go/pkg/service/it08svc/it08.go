package it08svc

import (
	"encoding/json"
	"test_project/pkg/model/it08model"
	"test_project/pkg/repository/it08repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IQuestionService interface {
	GetQuestions() (*[]it08model.QuestionResponse, error)
	InsertQuestion(req it08model.InsertQuestionRequest) error
	DeleteQuestion(id string) error
}

type QuestionService struct {
	QuestionRepository it08repository.IQuestionRepository
}

func NewQuestionService(dbcon *gorm.DB) IQuestionService {
	return &QuestionService{
		QuestionRepository: it08repository.NewQuestionRepository(dbcon),
	}
}

func (s *QuestionService) GetQuestions() (*[]it08model.QuestionResponse, error) {
	res, err := s.QuestionRepository.GetQuestions()
	if err != nil {
		return nil, err
	}

	questions := make([]it08model.QuestionResponse, 0)
	for _, question := range *res {
		var choices []string
		if err := json.Unmarshal(question.Choices, &choices); err != nil {
			return nil, err
		}

		questions = append(questions, it08model.QuestionResponse{
			ID:           question.ID,
			QuestionText: question.QuestionText,
			Choices:      choices,
			CreatedAt:    question.CreatedAt,
		})
	}

	return &questions, nil
}

func (s *QuestionService) InsertQuestion(req it08model.InsertQuestionRequest) error {

	order, err := s.QuestionRepository.GetNextSortOrder()
	if err != nil {
		return err
	}

	choicesJSON, err := json.Marshal(req.Choices)
	if err != nil {
		return err
	}

	question := it08model.Question{
		ID:           uuid.New().String(),
		QuestionText: req.QuestionText,
		Choices:      choicesJSON,
		SortOrder:    *order,
	}

	if err := s.QuestionRepository.InsertQuestion(question); err != nil {
		return err
	}

	return nil
}

func (s *QuestionService) DeleteQuestion(id string) error {
	if err := s.QuestionRepository.DeleteQuestion(id); err != nil {
		return err
	}

	reorderErr := s.QuestionRepository.Reorder()
	if reorderErr != nil {
		return reorderErr
	}

	return nil
}
