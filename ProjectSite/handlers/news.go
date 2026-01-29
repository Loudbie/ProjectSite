package handlers

import (
	"ProjectSite/News"
	"ProjectSite/database"

	"github.com/spf13/viper"

	"github.com/gofiber/fiber/v2"
)

var SecretKey = []byte(viper.GetString("secret_key"))

func GetNews(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT news.id, news.title, news.content, array_agg(Categories.name ORDER BY categories.name) AS Categories FROM news LEFT JOIN newscategories ON news.Id = newscategories.NewsId LEFT JOIN Categories ON newscategories.CategoryId = categories.Id GROUP BY news.id, news.Title, news.Content")
	if err != nil {
		return c.Status(500).SendString("Ошибка выполнения запроса к базе данных")
	}
	defer rows.Close()

	var result []News.GetResults
	for rows.Next() {
		var r News.GetResults
		err := rows.Scan(&r.ID, &r.Title, &r.Content, &r.Categories)
		if err != nil {
			return c.Status(500).SendString("Ошибка сканирования данных" + err.Error())
		}
		result = append(result, r)

	}

	return c.JSON(result)
}

func UpdateNews(c *fiber.Ctx) error {
	id := c.Params("id")
	news := new(News.News)

	if err := c.BodyParser(news); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	_, err := database.DB.Exec("UPDATE news SET title = $1, content = $2 WHERE news.id = $3",
		news.Title, news.Content, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления данных")
	}

	return c.SendString("Новость успешно обновлена")
}

func CreateNews(c *fiber.Ctx) error {
	news := new(News.News)
	if err := c.BodyParser(news); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	_, err := database.DB.Exec("INSERT INTO news (title, content) VALUES ($1, $2)",
		news.Title, news.Content)
	if err != nil {
		return c.Status(500).SendString("Ошибка вставки данных в базу")
	}

	return c.Status(201).SendString("Новость успешно создан")
}
