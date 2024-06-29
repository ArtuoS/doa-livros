package main

import (
	"strings"

	"github.com/ArtuoS/doa-livros/internal/controller"
	"github.com/ArtuoS/doa-livros/internal/entity"
	"github.com/ArtuoS/doa-livros/internal/repository"
	"github.com/ArtuoS/doa-livros/internal/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
)

func main() {
	server := server.SetupDatabase()

	bookRepo := &repository.BookRepository{DB: server.DB}
	userRepo := &repository.UserRepository{DB: server.DB}
	donatedBookRepo := &repository.DonatedBookRepository{DB: server.DB}

	userController := controller.NewUserController(userRepo)
	bookController := controller.NewBookController(bookRepo, donatedBookRepo)

	app := fiber.New(fiber.Config{
		Views: html.New("././web/views", ".html"),
	})

	app.Static("/", "././web/static")

	app.Get("/", bookController.GetAllBooks)
	app.Get("/books", bookController.GetAllBooks)

	app.Get("/users/auth", userController.GetAuthentication)
	app.Post("/users/auth", userController.Authenticate)

	app.Use(JWTMiddleware())

	app.Post("/books/redeem", bookController.RedeemBook)
	app.Get("/users/profile/:id", userController.GetUser)

	app.Listen("127.0.0.1:8080")
}

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		if tokenString == "" {
			return c.Redirect("/users/auth")
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims := &entity.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("KLw1agcGnuZbZvtRjG"), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.Redirect("/users/auth")
			}
			return c.Redirect("/users/auth")
		}

		if !token.Valid {
			return c.Redirect("/users/auth")
		}

		c.Locals("user", claims)
		return c.Next()
	}
}
