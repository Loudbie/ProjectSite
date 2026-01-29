package JWTToken

import (
	"ProjectSite/Auth"
	"fmt"
	"log"

	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Auth.UserData{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("херня какая-то: %v", token.Header["alg"])
		}
		log.Println("Вторая")
		return viper.GetString("secret_key"), nil
	})
	if err != nil {
		log.Println("Предпоследняя")
		log.Println(err)
		return nil, err
	}
	log.Println("Последняя")
	return token, nil
}

func Login(c *fiber.Ctx) error {
	// Создание маппы данных пользователя
	claims := Auth.CreateUserData(1, "Loudbie")

	// Создание токена с данными пользователя
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	tokenString, err := token.SignedString([]byte(viper.GetString("secret_key")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ошибка создания токена",
		})
	}
	log.Println("tokenstring: " + tokenString)
	log.Println(token)
	// Return the token
	return c.JSON(fiber.Map{
		"token":      tokenString,
		"expires_in": "6 часа",
		"user": fiber.Map{
			"id":       1,
			"username": "Loudbie",
		},
	})
}

func Authentication(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if auth == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "А где Bearer?",
		})
	}
	tokenString := strings.TrimPrefix(auth, "Bearer ")
	// Вырезаем нужную(!) хуйню(!), если вдруг вместо Bearer идёт Digest/Basic
	if tokenString == auth {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "С Bearer пиши, сука",
		})
	}
	log.Println("tokenstring: " + tokenString)

	token, err := parseToken(tokenString)

	log.Println(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Бля, ты кто?",
			"details": err.Error(),
		})
	}
	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Токен хуйня",
		})
	}
	if claims, ok := token.Claims.(*Auth.UserData); ok {
		c.Locals("ID", claims.ID)
		c.Locals("Name", claims.Name)
	}
	return c.Next()
}

func Protected(c *fiber.Ctx) error {
	userID := c.Locals("ID")
	userName := c.Locals("Name")

	return c.JSON(fiber.Map{
		"message":  "Проверка пройдена!",
		"user_id":  userID,
		"username": userName,
	})
}
