package it02handler

import (
	"crypto/rand"
	"test_project/pkg/model/it02model"
	"test_project/pkg/service/it02svc"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type IUserHandlers interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type UserHandler struct {
	UserService it02svc.IUserService
}

func NewUserHandler(dbcon *gorm.DB) IUserHandlers {
	return &UserHandler{
		UserService: it02svc.NewUserService(dbcon),
	}
}

func randomSecret() []byte {
	b := make([]byte, 32)
	_, _ = rand.Read(b)
	return b
}

var jwtSecret = randomSecret()

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req it02model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if req.Username == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Username and password required"})
	}

	user, err := h.UserService.Login(req)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.Status(200).JSON(fiber.Map{
		"token":    tokenString,
		"username": user.Username,
	})
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req it02model.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if req.Username == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Username and password required"})
	}

	if req.Password != req.ConfirmPassword {
		return c.Status(400).JSON(fiber.Map{"error": "Passwords do not match"})
	}

	err := h.UserService.Register(req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Register success"})
}
