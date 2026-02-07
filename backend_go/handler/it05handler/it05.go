package it05handler

import (
	"test_project/pkg/service/it05svc"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IQueueHandler interface {
	GetNextQueue(c *fiber.Ctx) error
	ClearQueue(c *fiber.Ctx) error
	GetCurrentQueue(c *fiber.Ctx) error
}

type QueueHandler struct {
	QueueService it05svc.IQueueService
}

func NewQueueHandler(dbcon *gorm.DB) IQueueHandler {
	return &QueueHandler{
		QueueService: it05svc.NewQueueService(dbcon),
	}
}

func (h *QueueHandler) GetNextQueue(c *fiber.Ctx) error {
	queue, err := h.QueueService.GetNextQueue()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(queue)
}

func (h *QueueHandler) ClearQueue(c *fiber.Ctx) error {
	queue, err := h.QueueService.ClearQueue()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(queue)
}

func (h *QueueHandler) GetCurrentQueue(c *fiber.Ctx) error {
	queue, err := h.QueueService.GetCurrentQueue()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(queue)
}
