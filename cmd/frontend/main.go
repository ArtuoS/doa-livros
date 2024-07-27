package main

import (
	"strings"

	"github.com/ArtuoS/doa-livros/entity"
	"github.com/ArtuoS/doa-livros/handler"
	"github.com/ArtuoS/doa-livros/infrastructure/repository"
	"github.com/ArtuoS/doa-livros/server"
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

	userHandler := handler.NewUserHandler(userRepo)
	bookHandler := handler.NewBookHandler(bookRepo, donatedBookRepo)

	app := fiber.New(fiber.Config{
		Views: html.New("././web/views", ".html"),
	})

	app.Static("/", "././web/static")

	app.Get("/auth", userHandler.GetAuthentication)
	app.Post("/auth", userHandler.Authenticate)

	app.Post("/books/redeem", bookHandler.RedeemBook)
	app.Get("/users/:id/profile", userHandler.GetUser)

	app.Get("/", bookHandler.GetAllBooks)
	app.Get("/books", bookHandler.GetAllBooks)
	app.Put("/books/:id/donate", bookHandler.AddBookToDonation)
	app.Post("/books", bookHandler.CreateBook)

	app.Listen("127.0.0.1:8080")

	// app.Use(JWTMiddleware())
}

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		if tokenString == "" {
			return c.Redirect("/auth")
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims := &entity.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("KLw1agcGnuZbZvtRjG"), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.Redirect("/auth")
			}
			return c.Redirect("/auth")
		}

		if !token.Valid {
			return c.Redirect("/auth")
		}

		c.Locals("user", claims)
		return c.Next()
	}
}
