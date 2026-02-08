package it10handler

import (
	"test_project/pkg/model/it10model"
	"test_project/pkg/service/it10svc"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IExamHandler interface {
	GetQuestions(c *fiber.Ctx) error
	SubmitExam(c *fiber.Ctx) error
}

type ExamHandler struct {
	ExamService it10svc.IExamService
}

func NewExamHandler(dbcon *gorm.DB) IExamHandler {
	return &ExamHandler{
		ExamService: it10svc.NewExamService(dbcon),
	}
}

func (h *ExamHandler) GetQuestions(c *fiber.Ctx) error {
	res, err := h.ExamService.GetQuestions()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(res)
}

func (h *ExamHandler) SubmitExam(c *fiber.Ctx) error {
	var req it10model.SubmitRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	if req.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "name is required",
		})
	}

	result, err := h.ExamService.SubmitExam(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(result)
}
