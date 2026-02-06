package it01svc

import (
	"test_project/pkg/model/it01model"
	"test_project/pkg/repository/it01repository"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IPersonServices interface {
	AddPerson(req it01model.Req) (*it01model.Person, error)
	ListPeople() (*[]it01model.Person, error)
	GetPerson(id string) (*it01model.Person, error)
}

type PersonServices struct {
	PersonRepository it01repository.IPersonRepository
}

func NewPersonService(dbcon *gorm.DB) IPersonServices {
	return &PersonServices{
		PersonRepository: it01repository.NewPersonRepository(dbcon),
	}
}

func (s *PersonServices) AddPerson(req it01model.Req) (*it01model.Person, error) {
	birthYear, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return nil, err
	}
	currentYear := time.Now().Year()
	age := currentYear - birthYear.Year()
	person := it01model.Person{
		ID:        uuid.New().String(),
		FullName:  req.FullName,
		BirthDate: req.BirthDate,
		Age:       age,
		Address:   req.Address,
	}
	err = s.PersonRepository.AddPerson(&person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (s *PersonServices) ListPeople() (*[]it01model.Person, error) {
	people, err := s.PersonRepository.ListPeople()
	if err != nil {
		return nil, err
	}
	return people, nil
}

func (s *PersonServices) GetPerson(personID string) (*it01model.Person, error) {
	person, err := s.PersonRepository.GetPerson(personID)
	if err != nil {
		return nil, err
	}
	return person, nil
}
