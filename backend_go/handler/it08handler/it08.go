package it08handler

import (
	"test_project/pkg/model/it08model"
	"test_project/pkg/service/it08svc"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IQuestionHandler interface {
	GetQuestions(c *fiber.Ctx) error
	InsertQuestion(c *fiber.Ctx) error
	DeleteQuestion(c *fiber.Ctx) error
}

type QuestionHandler struct {
	QuestinService it08svc.IQuestionService
}

func NewQuestionHandler(dbcon *gorm.DB) *QuestionHandler {
	return &QuestionHandler{
		QuestinService: it08svc.NewQuestionService(dbcon),
	}
}

func (h *QuestionHandler) GetQuestions(c *fiber.Ctx) error {
	questions, err := h.QuestinService.GetQuestions()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get questions"})
	}

	return c.Status(200).JSON(questions)
}

func (h *QuestionHandler) InsertQuestion(c *fiber.Ctx) error {
	var req it08model.InsertQuestionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if len(req.Choices) < 4 {
		return c.Status(400).JSON(fiber.Map{"error": "Required 4 choices"})
	}

	if err := h.QuestinService.InsertQuestion(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Question inserted successfully"})
}

func (h *QuestionHandler) DeleteQuestion(c *fiber.Ctx) error {
	questionID := c.Query("id")
	if questionID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Question ID is required"})
	}

	if err := h.QuestinService.DeleteQuestion(questionID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Question deleted successfully"})
}
