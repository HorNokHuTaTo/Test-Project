package it04svc

import (
	"errors"
	"regexp"
	"test_project/pkg/model/it04model"
	"test_project/pkg/repository/it04repository"
	"time"

	"gorm.io/gorm"
)

type IProfileService interface {
	InsertProfile(req *it04model.InsertProfileRequest) error
}

type ProfileService struct {
	ProfileRepository it04repository.IProfileRepository
}

func NewProfileService(dbconn *gorm.DB) IProfileService {
	return &ProfileService{
		ProfileRepository: it04repository.NewProfileRepository(dbconn),
	}
}

func (s *ProfileService) InsertProfile(req *it04model.InsertProfileRequest) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email) {
		return errors.New("invalid email format")
	}

	phoneRegex := regexp.MustCompile(`^0[0-9]{9}$`)
	if !phoneRegex.MatchString(req.Phone) {
		return errors.New("invalid phone format")
	}

	birthday, err := time.Parse("02/01/2006", req.Birthday)
	if err != nil {
		return errors.New("birthday must be format dd/mm/yyyy")
	}

	profile := it04model.Profile{
		ID:         req.ID,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Email:      req.Email,
		Phone:      req.Phone,
		Profile:    req.Profile,
		Birthday:   birthday,
		Occupation: req.Occupation,
	}

	if err := s.ProfileRepository.InsertProfile(&profile); err != nil {
		return err
	}
	return nil
}
