package it05svc

import (
	"errors"
	"sync"
	"test_project/pkg/model/it05model"
	"test_project/pkg/repository/it05repository"

	"gorm.io/gorm"
)

type IQueueService interface {
	GetNextQueue() (*it05model.QueueResponse, error)
	ClearQueue() (*string, error)
	GetCurrentQueue() (*it05model.Queue, error)
}

type QueueService struct {
	QueueRepository it05repository.IQueueRepository
	mu              sync.Mutex
}

func NewQueueService(dbconn *gorm.DB) IQueueService {
	return &QueueService{
		QueueRepository: it05repository.NewQueueRepository(dbconn),
	}
}

func (s *QueueService) GetNextQueue() (*it05model.QueueResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	queue, err := s.QueueRepository.GetCurrentQueue()
	if err != nil {
		return nil, err
	}

	if queue == nil {
		err := s.QueueRepository.InitializeQueue()
		if err != nil {
			return nil, err
		}
		queue, err = s.QueueRepository.GetCurrentQueue()
		if err != nil || queue == nil {
			return nil, errors.New("cannot initialize queue")
		}
	}

	res := it05model.QueueResponse{
		CurrentQueue: queue.CurrentQueue,
		Message:      "Get current queue success",
	}

	next := getNextQueueNumber(res.CurrentQueue)
	queue.CurrentQueue = next

	if err := s.QueueRepository.UpdateCurrentQueue(*queue); err != nil {
		return nil, err
	}

	return &res, nil
}

func getNextQueueNumber(current string) string {
	if len(current) < 2 {
		return "A0"
	}

	letter := current[0]
	digit := current[1]

	if digit < '9' {
		return string(letter) + string(digit+1)
	}

	if letter < 'Z' {
		return string(letter+1) + "0"
	}

	return "A0"
}

func getPreviousQueue(current string) string {
	if len(current) < 2 {
		return "A0"
	}
	letter := current[0]
	digit := current[1]
	if digit > '0' {
		return string(letter) + string(digit-1)
	}
	if letter > 'A' {
		return string(letter-1) + "9"
	}
	return "A0"
}

func (s *QueueService) ClearQueue() (*string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	queue, err := s.QueueRepository.GetCurrentQueue()
	if err != nil {
		return nil, err
	}

	if queue == nil {
		return nil, errors.New("queue not found")
	}

	queue.CurrentQueue = "A0"

	if err := s.QueueRepository.UpdateCurrentQueue(*queue); err != nil {
		return nil, err
	}

	resetQueue := "A0"
	res := "Queue cleared, reset to " + resetQueue

	return &res, nil
}

func (s *QueueService) GetCurrentQueue() (*it05model.Queue, error) {
	queue, err := s.QueueRepository.GetCurrentQueue()
	if err != nil {
		return nil, err
	}

	currentQueue := getPreviousQueue(queue.CurrentQueue)
	queue.CurrentQueue = currentQueue

	return queue, nil
}
