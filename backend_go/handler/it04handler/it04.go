package it04handler

import (
	"test_project/pkg/model/it04model"
	"test_project/pkg/service/it04svc"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type IProfileHandler interface {
	InsertProfile(c *fiber.Ctx) error
	GetOccupations(c *fiber.Ctx) error
}

type ProfileHandler struct {
	ProfileService it04svc.IProfileService
}

func NewProfileHandler(dbcon *gorm.DB) *ProfileHandler {
	return &ProfileHandler{
		ProfileService: it04svc.NewProfileService(dbcon),
	}
}

func (h *ProfileHandler) InsertProfile(c *fiber.Ctx) error {
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")
	email := c.FormValue("email")
	phone := c.FormValue("phone")
	birthday := c.FormValue("birthday")
	occupation := c.FormValue("occupation")
	sex := c.FormValue("sex")

	for _, field := range []string{firstName, lastName, email, phone, birthday, occupation, sex} {
		if field == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "All fields are required",
			})
		}
	}

	file, err := c.FormFile("profile")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Profile file is required",
		})
	}

	req := it04model.InsertProfileRequest{
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Phone:      phone,
		Birthday:   birthday,
		Occupation: occupation,
		Sex:        sex,
		Profile:    file,
	}

	id, err := h.ProfileService.InsertProfile(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Save Data Success",
		"id":      id,
	})
}

func (h *ProfileHandler) GetOccupations(c *fiber.Ctx) error {
	res, err := h.ProfileService.GetOccupations()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve occupations",
		})
	}

	return c.Status(200).JSON(res)
}
