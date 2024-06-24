package main

import (
	"log"
	"net/http"

	"github.com/ArtuoS/doa-livros/internal/entity"
	"github.com/ArtuoS/doa-livros/internal/repository"
	"github.com/ArtuoS/doa-livros/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Books(c *fiber.Ctx) error {
	bookRepo := repository.BookRepository{DB: DB}
	books, err := bookRepo.GetAllBooks()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Render("books", fiber.Map{"Books": books})
}

func Redeem(c *fiber.Ctx) error {
	var donatedBook entity.DonatedBook
	if err := c.BodyParser(&donatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	model := usecase.NewDonateBookModel(&donatedBook)
	if err := usecase.NewDonateBookUseCase(c.BookRepo, c.DonatedBookRepo, model).Handle(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return c.SendString("Redeem request processed successfully")
}

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres password=root dbname=doalivros sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	DB = db

	app := fiber.New(fiber.Config{
		Views: html.New("././web/views", ".html"),
	})

	app.Static("/", "././web/static")

	app.Get("/books", Books)
	app.Post("/books/redeem", Redeem)

	app.Listen(":8080")
}
