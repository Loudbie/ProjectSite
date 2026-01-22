package main

import (
	"ProjectSite/routes"
	"log"

	"ProjectSite/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	// инициализируем базу данных
	if err := database.Connect(); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// создаём новое приложение Fiber
	app := fiber.New(fiber.Config{
		Prefork: true, // используем предварительное форкование для увеличения производительности
	})

	// Подключаем middleware
	app.Use(logger.New())   // Логирование запросов
	app.Use(compress.New()) // Сжатие ответов
	app.Use(recover.New())  // Восстановление после паники
	app.Use(limiter.New())  // Лимит запросов для предотвращения DDOS атак

	// Регистрация маршрутов

	routes.RegisterNewsRoutes(app)

	// Запускаем сервер
	log.Fatal(app.Listen(":3000"))
}
