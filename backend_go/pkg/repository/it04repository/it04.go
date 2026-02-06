package it04repository

import (
	"context"
	"test_project/pkg/model/it04model"
	"time"

	"gorm.io/gorm"
)

type IProfileRepository interface {
	InsertProfile(profile *it04model.Profile) error
}

type ProfileRepository struct {
	DB *gorm.DB
}

func NewProfileRepository(dbconn *gorm.DB) IProfileRepository {
	return &ProfileRepository{
		DB: dbconn,
	}
}

func (r *ProfileRepository) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func (r *ProfileRepository) InsertProfile(profile *it04model.Profile) error {
	ctx, cancel := r.withTimeout()
	defer cancel()

	err := r.DB.Table("it04").WithContext(ctx).Create(profile).Error
	if err != nil {
		return err
	}

	return nil
}
