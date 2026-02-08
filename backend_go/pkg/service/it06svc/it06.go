package it06svc

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image/png"
	"regexp"
	"test_project/pkg/model/it06model"
	"test_project/pkg/repository/it06repository"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code39"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ISkuService interface {
	GetSkus() (*[]it06model.Sku, error)
	InsertSku(req it06model.InsertSkuRequest) error
	DeleteSku(id string) error
}

type SkuService struct {
	SkuRepository it06repository.ISkuRepository
}

func NewSkuService(dbconn *gorm.DB) ISkuService {
	return &SkuService{
		SkuRepository: it06repository.NewSkuRepository(dbconn),
	}
}

func (s *SkuService) GetSkus() (*[]it06model.Sku, error) {
	res, err := s.SkuRepository.GetSkus()
	if err != nil {
		return nil, err
	}

	skus := make([]it06model.Sku, 0)
	for _, sku := range *res {
		skus = append(skus, it06model.Sku{
			ID:        sku.ID,
			SkuCode:   sku.SkuCode,
			Barcode:   sku.Barcode,
			CreatedAt: sku.CreatedAt,
		})
	}

	return &skus, nil
}

func (s *SkuService) InsertSku(req it06model.InsertSkuRequest) error {
	validateSkuCode := regexp.MustCompile(`^[A-Z0-9]{4}(?:-[A-Z0-9]{4}){3}$`)
	if !validateSkuCode.MatchString(req.SkuCode) {
		return errors.New("invalid sku code format")
	}

	barcode, err := GenerateCode39Barcode(req.SkuCode)
	if err != nil {
		return err
	}

	newSku := it06model.Sku{
		ID:      uuid.New().String(),
		SkuCode: req.SkuCode,
		Barcode: *barcode,
	}

	if err := s.SkuRepository.InsertSku(newSku); err != nil {
		return err
	}

	return nil
}

func GenerateCode39Barcode(sku string) (*string, error) {
	var bc barcode.Barcode // 👈 บังคับ type เป็น interface

	var err error
	bc, err = code39.Encode(sku, true, false)
	if err != nil {
		return nil, err
	}

	bc, err = barcode.Scale(bc, 400, 100)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, bc); err != nil {
		return nil, err
	}

	result := base64.StdEncoding.EncodeToString(buf.Bytes())
	return &result, nil
}

func (s *SkuService) DeleteSku(id string) error {
	if err := s.SkuRepository.DeleteSku(id); err != nil {
		return err
	}

	return nil
}
