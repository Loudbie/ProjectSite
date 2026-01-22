package handlers

import (
	"ProjectSite/News"
	"ProjectSite/database"

	"github.com/gofiber/fiber/v2"
)

func GetNews(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT id, title, content FROM news")
	if err != nil {
		return c.Status(500).SendString("Ошибка выполнения запроса к базе данных")
	}
	defer rows.Close()

	var news []News.News
	for rows.Next() {
		var new News.News
		err := rows.Scan(&new.ID, &new.Title, &new.Content)
		if err != nil {
			return c.Status(500).SendString("Ошибка сканирования данных")
		}
		news = append(news, new)
	}

	return c.JSON(news)
}

// обновление новости
func UpdateNews(c *fiber.Ctx) error {
	id := c.Params("id")
	news := new(News.News)

	if err := c.BodyParser(news); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	_, err := database.DB.Exec("UPDATE news SET title = $1, content = $2, WHERE id = $3",
		news.Title, news.Content, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления данных")
	}

	return c.SendString("Новость успешно обновлена")
}
