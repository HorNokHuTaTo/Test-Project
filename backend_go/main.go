package main

import (
	"fmt"
	"test_project/handler/it01handler"
	"test_project/handler/it02handler"
	"test_project/handler/it03handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupITRoutes(api fiber.Router, db *gorm.DB) {
	it01Handler := it01handler.NewPersonHandler(db)
	it02Handler := it02handler.NewUserHandler(db)
	it01Group := api.Group("it01")
	it01Group.
		Post("/people", it01Handler.AddPerson).
		Get("/people", it01Handler.ListPeople).
		Get("/people/:id", it01Handler.GetPerson)

	it02Group := api.Group("it02")
	it02Group.
		Post("/login", it02Handler.Login).
		Post("/register", it02Handler.Register)

	it03Group := api.Group("it03")
	it03Handler := it03handler.NewDocumentHandler(db)
	it03Group.
		Patch("/update-status", it03Handler.UpdateDocumentsStatus).
		Get("/documents", it03Handler.GetAllDocuments)
}

func main() {
	// change user password base on db
	dsn := "root:Bright1998@@(localhost:3306)/exam?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Check connection
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get sql.DB")
	}
	if err := sqlDB.Ping(); err != nil {
		panic("failed to ping database: " + err.Error())
	}
	fmt.Println("Database connection successful!")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api")
	SetupITRoutes(api, db)

	app.Listen(":3000")
}
