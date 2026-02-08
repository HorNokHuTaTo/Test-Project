package it06handler

import (
	"test_project/pkg/model/it06model"
	"test_project/pkg/service/it06svc"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ISkuHandler interface {
	GetSkus(c *fiber.Ctx) error
	InsertSku(c *fiber.Ctx) error
	DeleteSku(c *fiber.Ctx) error
}

type SkuHandler struct {
	SkuService it06svc.ISkuService
}

func NewSkuHandler(dbconn *gorm.DB) ISkuHandler {
	return &SkuHandler{
		SkuService: it06svc.NewSkuService(dbconn),
	}
}

func (h *SkuHandler) GetSkus(c *fiber.Ctx) error {
	skus, err := h.SkuService.GetSkus()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get SKUs"})
	}

	return c.Status(200).JSON(skus)
}

func (h *SkuHandler) InsertSku(c *fiber.Ctx) error {
	var req it06model.InsertSkuRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.SkuService.InsertSku(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "SKU inserted successfully"})
}

func (h *SkuHandler) DeleteSku(c *fiber.Ctx) error {
	skuID := c.Query("id")
	if skuID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "SKU ID is required"})
	}

	if err := h.SkuService.DeleteSku(skuID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "SKU deleted successfully"})
}
