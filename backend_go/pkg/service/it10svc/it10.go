package it10svc

import (
	"test_project/pkg/model/it10model"
	"test_project/pkg/repository/it10repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IExamService interface {
	GetQuestions() (*[]it10model.Question, error)
	SubmitExam(req it10model.SubmitRequest) (*it10model.ExamResult, error)
}

type ExamService struct {
	ExamRepository it10repository.IExamRepository
}

func NewExamService(dbcon *gorm.DB) IExamService {
	return &ExamService{
		ExamRepository: it10repository.NewExamRepository(dbcon),
	}
}

func (s *ExamService) GetQuestions() (*[]it10model.Question, error) {
	question := []it10model.Question{
		{
			QuestionText:  "1. ข้อใดคือจำนวนคี่",
			Choices:       []string{"3", "5", "9", "11"},
			CorrectChoice: 1,
		},
		{
			QuestionText:  "2. X + 2 = 4 จงหา X",
			Choices:       []string{"1", "2", "3", "4"},
			CorrectChoice: 2,
		},
	}

	return &question, nil
}

func (s *ExamService) SubmitExam(req it10model.SubmitRequest) (*it10model.ExamResult, error) {
	questions, err := s.GetQuestions()
	if err != nil {
		return nil, err
	}

	score := 0
	total := len(*questions)

	for i, q := range *questions {
		if i < len(req.Answers) && req.Answers[i] == q.CorrectChoice {
			score++
		}
	}

	result := it10model.ExamResult{
		ID:    uuid.New().String(),
		Name:  req.Name,
		Score: score,
		Total: total,
	}

	if err := s.ExamRepository.SubmitExam(result); err != nil {
		return nil, err
	}

	return &result, nil
}
