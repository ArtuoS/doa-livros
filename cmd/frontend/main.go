package main

import (
	"log"
	"strings"

	"github.com/ArtuoS/doa-livros/internal/controller"
	"github.com/ArtuoS/doa-livros/internal/entity"
	"github.com/ArtuoS/doa-livros/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres password=root dbname=doalivros sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	bookRepo := &repository.BookRepository{DB: db}
	userRepo := &repository.UserRepository{DB: db}
	donatedBookRepo := &repository.DonatedBookRepository{DB: db}

	userController := controller.NewUserController(userRepo)
	bookController := controller.NewBookController(bookRepo, donatedBookRepo)

	DB = db

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
