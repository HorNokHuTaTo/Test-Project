package it02repository

import (
	"context"
	"test_project/pkg/model/it02model"
	"time"

	"gorm.io/gorm"
)

type IUserRepository interface {
	RegisterUser(user it02model.User) error
	GetUserByUsername(username string) (*it02model.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(dbconn *gorm.DB) IUserRepository {
	return &UserRepository{
		DB: dbconn,
	}
}

func (r *UserRepository) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func (r *UserRepository) RegisterUser(user it02model.User) error {
	ctx, cancel := r.withTimeout()
	defer cancel()
	err := r.DB.Table("it02").WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(username string) (*it02model.User, error) {
	ctx, cancel := r.withTimeout()
	defer cancel()
	var user it02model.User
	err := r.DB.Table("it02").WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
