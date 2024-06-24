package main

import (
	"log"

	"github.com/ArtuoS/doa-livros/internal/controller"
	"github.com/ArtuoS/doa-livros/internal/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
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

	app.Get("/books", bookController.GetAllBooks)
	app.Post("/books/redeem", bookController.RedeemBook)

	app.Get("/users/profile/:id", userController.GetUser)

	app.Listen(":8080")
}
