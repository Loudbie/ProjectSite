package routes

import (
	"ProjectSite/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterNewsRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/news", handlers.GetNews)        // Получить все новости
	api.Put("/news/:id", handlers.UpdateNews) // Обновить новости
}
