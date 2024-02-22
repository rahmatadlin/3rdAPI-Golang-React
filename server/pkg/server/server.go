package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rahmatadlin/3rdAPI-Golang-React/pkg/controller"
)

func AppWithRoutes() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Routes for 3rd party API integration
	app.Get("/api/books", controller.ReadBooksHandler)

	return app
}
