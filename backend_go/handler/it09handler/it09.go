package it09handler

import (
	"test_project/pkg/model/it09model"
	"test_project/pkg/service/it09svc"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ICommentHandler interface {
	GetComments(c *fiber.Ctx) error
	InsertComment(c *fiber.Ctx) error
}

type CommentHandler struct {
	CommentService it09svc.ICommentService
}

func NewCommentHandler(dbcon *gorm.DB) ICommentHandler {
	return &CommentHandler{
		CommentService: it09svc.NewCommentService(dbcon),
	}
}

func (h *CommentHandler) GetComments(c *fiber.Ctx) error {
	comments, err := h.CommentService.GetComments()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot get comments",
		})
	}

	return c.Status(200).JSON(comments)
}

func (h *CommentHandler) InsertComment(c *fiber.Ctx) error {
	var req it09model.InsertCommentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	if req.Message == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "message is required",
		})
	}

	if err := h.CommentService.InsertComment(req); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "cannot insert comment",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "comment inserted successfully",
	})
}
