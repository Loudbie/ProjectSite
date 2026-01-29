package routes

import (
	"ProjectSite/JWTToken"
	"ProjectSite/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterNewsRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/login", JWTToken.Login)                                 // Создание Токена
	api.Get("/protected", JWTToken.Authentication, JWTToken.Protected) //Защищённый маршрут
	api.Get("/news", handlers.GetNews)                                 // Получить все новости
	api.Put("/news/:id", handlers.UpdateNews)                          // Обновить новости
	api.Post("/news/:id", handlers.CreateNews)                         // Добавить новость
}
