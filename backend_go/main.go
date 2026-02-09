package main

import (
	"fmt"
	"test_project/handler/it01handler"
	"test_project/handler/it02handler"
	"test_project/handler/it03handler"
	"test_project/handler/it04handler"
	"test_project/handler/it05handler"
	"test_project/handler/it06handler"
	"test_project/handler/it07handler"
	"test_project/handler/it08handler"
	"test_project/handler/it09handler"
	"test_project/handler/it10handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupITRoutes(api fiber.Router, db *gorm.DB) {
	it01Handler := it01handler.NewPersonHandler(db)
	it01Group := api.Group("it01")
	it01Group.
		Post("/people", it01Handler.AddPerson).
		Get("/people", it01Handler.ListPeople).
		Get("/people/:id", it01Handler.GetPerson)

	it02Handler := it02handler.NewUserHandler(db)
	it02Group := api.Group("it02")
	it02Group.
		Post("/login", it02Handler.Login).
		Post("/register", it02Handler.Register)

	it03Handler := it03handler.NewDocumentHandler(db)
	it03Group := api.Group("it03")
	it03Group.
		Patch("/update-status", it03Handler.UpdateDocumentsStatus).
		Get("/documents", it03Handler.GetAllDocuments)

	it04Handler := it04handler.NewProfileHandler(db)
	it04Group := api.Group("it04")
	it04Group.
		Post("/insert-profile", it04Handler.InsertProfile).
		Get("/occupations", it04Handler.GetOccupations)

	it05Handler := it05handler.NewQueueHandler(db)
	it05Group := api.Group("it05")
	it05Group.
		Post("/next-queue", it05Handler.GetNextQueue).
		Post("/clear-queue", it05Handler.ClearQueue).
		Get("/current-queue", it05Handler.GetCurrentQueue)

	it06Handler := it06handler.NewSkuHandler(db)
	it06Group := api.Group("it06")
	it06Group.
		Get("/skus", it06Handler.GetSkus).
		Post("/insert-sku", it06Handler.InsertSku).
		Delete("/delete-sku", it06Handler.DeleteSku)

	it07Handler := it07handler.NewSku2Handler(db)
	it07Group := api.Group("it07")
	it07Group.
		Get("/skus", it07Handler.GetSkus).
		Post("/insert-sku", it07Handler.InsertSku).
		Delete("/delete-sku", it07Handler.DeleteSku)

	it08Handler := it08handler.NewQuestionHandler(db)
	it08Group := api.Group("it08")
	it08Group.
		Get("/questions", it08Handler.GetQuestions).
		Post("/insert-question", it08Handler.InsertQuestion).
		Delete("/delete-question", it08Handler.DeleteQuestion)

	it09Handler := it09handler.NewCommentHandler(db)
	it09Group := api.Group("it09")
	it09Group.
		Get("/comments", it09Handler.GetComments).
		Post("/insert-comment", it09Handler.InsertComment)

	it10Handler := it10handler.NewExamHandler(db)
	it10Group := api.Group("it10")
	it10Group.
		Get("/questions", it10Handler.GetQuestions).
		Post("/submit-exam", it10Handler.SubmitExam)
}

func main() {
	// change user password base on db
	dsn := "root:LsCMsrPvzMYeaaPuaqpRXOQuODRFHEHu@(maglev.proxy.rlwy.net:16055)/exam?charset=utf8mb4&parseTime=True&loc=Local"
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
