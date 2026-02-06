package it02svc

import (
	"fmt"
	"test_project/pkg/model/it02model"
	"test_project/pkg/repository/it02repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserService interface {
	Login(req it02model.LoginRequest) (*it02model.User, error)
	Register(req it02model.RegisterRequest) error
}

type UserService struct {
	UserRepository it02repository.IUserRepository
}

func NewUserService(dbcon *gorm.DB) IUserService {
	return &UserService{
		UserRepository: it02repository.NewUserRepository(dbcon),
	}
}

func (s *UserService) Login(req it02model.LoginRequest) (*it02model.User, error) {
	user, err := s.UserRepository.GetUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Register(req it02model.RegisterRequest) error {
	user, err := s.UserRepository.GetUserByUsername(req.Username)
	if err == nil && user != nil {
		return fmt.Errorf("username already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := it02model.User{
		ID:       uuid.New().String(),
		Username: req.Username,
		Password: string(hash),
	}

	if err := s.UserRepository.RegisterUser(newUser); err != nil {
		return err
	}

	return nil
}
