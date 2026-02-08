package it03handler

import (
	"test_project/pkg/model/it03model"
	"test_project/pkg/service/it03svc"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IDocumentHandlers interface {
	UpdateDocumentsStatus(c *fiber.Ctx) error
	GetAllDocuments(c *fiber.Ctx) error
}

type DocumentHandler struct {
	DocumentService it03svc.IDocumentService
}

func NewDocumentHandler(dbcon *gorm.DB) *DocumentHandler {
	return &DocumentHandler{
		DocumentService: it03svc.NewDocumentService(dbcon),
	}
}

func (h *DocumentHandler) UpdateDocumentsStatus(c *fiber.Ctx) error {
	var req it03model.UpdateStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if len(req.IDs) == 0 || req.Status == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Missing IDs or status",
		})
	}

	if !it03model.IsValidStatus(req.Status) {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid status",
		})
	}

	if err := h.DocumentService.UpdateDocumentsStatus(req); err != nil {
		if err.Error() == "ไม่สามารถอัปเดตรายการที่อนุมัติแล้วได้" {
			return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"error": "Update failed",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})
}

func (h *DocumentHandler) GetAllDocuments(c *fiber.Ctx) error {
	documents, err := h.DocumentService.GetAllDocuments()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve documents",
		})
	}

	return c.Status(200).JSON(documents)
}
