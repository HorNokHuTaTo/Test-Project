package it01repository

import (
	"context"
	"test_project/pkg/model/it01model"
	"time"

	"gorm.io/gorm"
)

type IPersonRepository interface {
	AddPerson(person *it01model.Person) error
	ListPeople() (*[]it01model.Person, error)
	GetPerson(id string) (*it01model.Person, error)
}

type PersonRepository struct {
	DB *gorm.DB
}

func NewPersonRepository(dbconn *gorm.DB) IPersonRepository {
	return &PersonRepository{
		DB: dbconn,
	}
}

func (r *PersonRepository) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func (r *PersonRepository) AddPerson(person *it01model.Person) error {
	ctx, cancel := r.withTimeout()
	defer cancel()
	err := r.DB.Table("it01").WithContext(ctx).Create(person).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *PersonRepository) ListPeople() (*[]it01model.Person, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var people []it01model.Person
	err := r.DB.Table("it01").WithContext(ctx).Find(&people).Error
	if err != nil {
		return nil, err
	}

	return &people, nil
}

func (r *PersonRepository) GetPerson(id string) (*it01model.Person, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()

	var person it01model.Person
	err := r.DB.Table("it01").WithContext(ctx).First(&person, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &person, nil
}
