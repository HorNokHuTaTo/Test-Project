package it07handler

import (
	"test_project/pkg/model/it07model"
	"test_project/pkg/service/it07svc"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ISku2Handler interface {
	GetSkus(c *fiber.Ctx) error
	InsertSku(c *fiber.Ctx) error
	DeleteSku(c *fiber.Ctx) error
}

type Sku2Handler struct {
	Sku2Service it07svc.ISku2Service
}

func NewSku2Handler(dbconn *gorm.DB) ISku2Handler {
	return &Sku2Handler{
		Sku2Service: it07svc.NewSku2Service(dbconn),
	}
}

func (h *Sku2Handler) GetSkus(c *fiber.Ctx) error {
	skus, err := h.Sku2Service.GetSkus()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get SKUs"})
	}

	return c.Status(200).JSON(skus)
}

func (h *Sku2Handler) InsertSku(c *fiber.Ctx) error {
	var req it07model.InsertSkuRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.Sku2Service.InsertSku(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "SKU inserted successfully"})
}

func (h *Sku2Handler) DeleteSku(c *fiber.Ctx) error {
	skuID := c.Query("id")
	if skuID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "SKU ID is required"})
	}

	if err := h.Sku2Service.DeleteSku(skuID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "SKU deleted successfully"})
}
