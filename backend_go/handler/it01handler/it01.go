package it01handler

import (
	"test_project/pkg/model/it01model"
	"test_project/pkg/service/it01svc"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IPersonHandlers interface {
	AddPerson(c *fiber.Ctx) error
	ListPeople(c *fiber.Ctx) error
	GetPerson(c *fiber.Ctx) error
}

type PersonHandler struct {
	PersonService it01svc.IPersonServices
}

func NewPersonHandler(dbcon *gorm.DB) IPersonHandlers {
	return &PersonHandler{
		PersonService: it01svc.NewPersonService(dbcon),
	}
}

func (h *PersonHandler) AddPerson(c *fiber.Ctx) error {
	var req it01model.Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if req.FullName == "" || req.BirthDate == "" || req.Address == "" {
		return c.Status(400).JSON(fiber.Map{"error": "All fields are required"})
	}

	person, err := h.PersonService.AddPerson(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid birth date format"})
	}

	return c.Status(201).JSON(person)
}

func (h *PersonHandler) ListPeople(c *fiber.Ctx) error {
	people, err := h.PersonService.ListPeople()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to list people"})
	}

	return c.Status(200).JSON(people)
}

func (h *PersonHandler) GetPerson(c *fiber.Ctx) error {
	personID := c.Params("id")
	if personID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid person ID"})
	}

	person, err := h.PersonService.GetPerson(personID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Person not found"})
	}

	return c.Status(200).JSON(person)
}
