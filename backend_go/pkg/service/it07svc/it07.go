package it07svc

import (
	"encoding/base64"
	"errors"
	"regexp"
	"test_project/pkg/model/it07model"
	"test_project/pkg/repository/it07repository"

	"github.com/google/uuid"
	qrcode "github.com/skip2/go-qrcode"
	"gorm.io/gorm"
)

type ISku2Service interface {
	GetSkus() (*[]it07model.Sku, error)
	InsertSku(req it07model.InsertSkuRequest) error
	DeleteSku(id string) error
}

type Sku2Service struct {
	Sku2Repository it07repository.ISku2Repository
}

func NewSku2Service(dbconn *gorm.DB) ISku2Service {
	return &Sku2Service{
		Sku2Repository: it07repository.NewSku2Repository(dbconn),
	}
}

func (s *Sku2Service) GetSkus() (*[]it07model.Sku, error) {
	res, err := s.Sku2Repository.GetSkus()
	if err != nil {
		return nil, err
	}

	skus := make([]it07model.Sku, 0)
	for _, sku := range *res {
		skus = append(skus, it07model.Sku{
			ID:        sku.ID,
			SkuCode:   sku.SkuCode,
			QrCode:    sku.QrCode,
			CreatedAt: sku.CreatedAt,
		})
	}

	return &skus, nil
}

func (s *Sku2Service) InsertSku(req it07model.InsertSkuRequest) error {
	validateSkuCode := regexp.MustCompile(`^[A-Z0-9]{5}(?:-[A-Z0-9]{5}){5}$`)
	if !validateSkuCode.MatchString(req.SkuCode) {
		return errors.New("invalid sku code format")
	}

	qrCode, err := GenerateQRCodeBase64(req.SkuCode)
	if err != nil {
		return err
	}

	newSku := it07model.Sku{
		ID:      uuid.New().String(),
		SkuCode: req.SkuCode,
		QrCode:  *qrCode,
	}

	if err := s.Sku2Repository.InsertSku(newSku); err != nil {
		return err
	}

	return nil
}

func GenerateQRCodeBase64(text string) (*string, error) {
	png, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	result := base64.StdEncoding.EncodeToString(png)
	return &result, nil
}

func (s *Sku2Service) DeleteSku(id string) error {
	if err := s.Sku2Repository.DeleteSku(id); err != nil {
		return err
	}

	return nil
}
