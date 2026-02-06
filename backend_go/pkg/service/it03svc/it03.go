package it03svc

import (
	"test_project/pkg/model/it03model"
	"test_project/pkg/repository/it03repository"

	"gorm.io/gorm"
)

type IDocumentService interface {
	UpdateDocumentsStatus(req it03model.UpdateStatusRequest) error
	GetAllDocuments() (*[]it03model.GetDocumentsResponse, error)
}

type DocumentService struct {
	DocumentRepository it03repository.IDocumentRepository
}

func NewDocumentService(dbcon *gorm.DB) IDocumentService {
	return &DocumentService{
		DocumentRepository: it03repository.NewDocumentRepository(dbcon),
	}
}

func (s *DocumentService) UpdateDocumentsStatus(req it03model.UpdateStatusRequest) error {
	err := s.DocumentRepository.UpdateDocumentsStatus(req)
	if err != nil {
		if err.Error() == "ไม่สามารถอัปเดตเอกสารที่อนุมัติแล้วได้" {
			return err
		}
		return err
	}
	return nil
}

func (s *DocumentService) GetAllDocuments() (*[]it03model.GetDocumentsResponse, error) {
	resp, err := s.DocumentRepository.GetAllDocuments()
	if err != nil {
		return nil, err
	}

	var documents []it03model.GetDocumentsResponse
	for _, doc := range *resp {
		documents = append(documents, it03model.GetDocumentsResponse{
			ID:          doc.ID,
			Title:       doc.Title,
			Description: doc.Description,
			Status:      doc.Status,
		})
	}

	return &documents, nil
}
