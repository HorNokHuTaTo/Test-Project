package it04svc

import (
	"encoding/base64"
	"errors"
	"io"
	"regexp"
	"test_project/pkg/model/it04model"
	"test_project/pkg/repository/it04repository"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IProfileService interface {
	InsertProfile(req it04model.InsertProfileRequest) (*string, error)
	GetOccupations() (*[]it04model.Occupation, error)
}

type ProfileService struct {
	ProfileRepository it04repository.IProfileRepository
}

func NewProfileService(dbconn *gorm.DB) IProfileService {
	return &ProfileService{
		ProfileRepository: it04repository.NewProfileRepository(dbconn),
	}
}

func (s *ProfileService) InsertProfile(req it04model.InsertProfileRequest) (*string, error) {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(req.Email) {
		return nil, errors.New("invalid email format")
	}

	phoneRegex := regexp.MustCompile(`^0[0-9]{9}$`)
	if !phoneRegex.MatchString(req.Phone) {
		return nil, errors.New("invalid phone format")
	}

	birthday, err := time.Parse("02/01/2006", req.Birthday)
	if err != nil {
		return nil, errors.New("birthday must be format dd/mm/yyyy")
	}

	occList, err := s.GetOccupations()
	if err != nil {
		return nil, err
	}

	selectedOccupation := ""
	for _, occ := range *occList {
		if occ.ID == req.Occupation {
			selectedOccupation = occ.Value
			break
		}
	}

	if selectedOccupation == "" {
		return nil, errors.New("invalid occupation")
	}

	var profileBase64 string
	if req.Profile != nil {
		f, err := req.Profile.Open()
		if err != nil {
			return nil, errors.New("failed to open profile file")
		}
		defer f.Close()

		fileBytes, err := io.ReadAll(f)
		if err != nil {
			return nil, errors.New("failed to read profile file")
		}
		profileBase64 = base64.StdEncoding.EncodeToString(fileBytes)
	}

	profile := it04model.Profile{
		ID:         uuid.New().String(),
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Email:      req.Email,
		Phone:      req.Phone,
		Profile:    profileBase64,
		Birthday:   birthday,
		Occupation: selectedOccupation,
		Sex:        req.Sex,
	}

	if err := s.ProfileRepository.InsertProfile(&profile); err != nil {
		return nil, err
	}
	return &profile.ID, nil
}

func (s *ProfileService) GetOccupations() (*[]it04model.Occupation, error) {
	Occupations := []it04model.Occupation{
		{ID: "1", Label: "Software Developer", Value: "Software Developer"},
		{ID: "2", Label: "UI/UX Designer", Value: "UI/UX Designer"},
		{ID: "3", Label: "QA / Tester", Value: "QA / Tester"},
		{ID: "4", Label: "Project Manager", Value: "Project Manager"},
		{ID: "5", Label: "System Analyst", Value: "System Analyst"},
		{ID: "6", Label: "DevOps Engineer", Value: "DevOps Engineer"},
	}

	if len(Occupations) == 0 {
		return nil, errors.New("no occupations found")
	}

	return &Occupations, nil
}
