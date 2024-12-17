package main

import (
	"github.com/ArtuoS/doa-livros/handler"
	"github.com/ArtuoS/doa-livros/infrastructure/repository"
	"github.com/ArtuoS/doa-livros/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
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
	app.Post("/users", userHandler.CreateUser)

	app.Get("/", bookHandler.GetAllBooks)
	app.Get("/books", bookHandler.GetAllBooks)
	app.Put("/books/:id/donate", bookHandler.AddBookToDonation)
	app.Post("/books", bookHandler.CreateBook)
	app.Delete("/books/:id", bookHandler.DeleteBook)

	app.Listen("127.0.0.1:8080")
}
