package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ArtuoS/doa-livros/entity"
	"github.com/ArtuoS/doa-livros/infrastructure/repository"
	"github.com/ArtuoS/doa-livros/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
)

type BookHandler struct {
	BookRepo        *repository.BookRepository
	DonatedBookRepo *repository.DonatedBookRepository
}

func NewBookHandler(bookRepo *repository.BookRepository, donatedBookRepo *repository.DonatedBookRepository) *BookHandler {
	return &BookHandler{
		BookRepo:        bookRepo,
		DonatedBookRepo: donatedBookRepo,
	}
}

func (b *BookHandler) CreateBook(ctx *fiber.Ctx) error {
	var book entity.Book
	if err := ctx.BodyParser(&book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	book.Donating = false
	if err := b.BookRepo.CreateBook(&book); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Redirect("/")
}

func (b *BookHandler) DeleteBook(ctx *fiber.Ctx) error {
	bookID, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	if err := b.BookRepo.DeleteBook(bookID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Redirect("/")
}

func (b *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "invalid book ID", http.StatusBadRequest)
		return
	}
	book, err := b.BookRepo.GetBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (b *BookHandler) GetAllBooks(c *fiber.Ctx) error {
	books, err := b.BookRepo.GetAllBooks()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Render("books", fiber.Map{"Books": books})
}

func (b *BookHandler) AddBookToDonation(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = b.BookRepo.AddBookToDonation(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	books, _ := b.BookRepo.GetAllBooks()
	return c.Render("books", fiber.Map{"Books": books})
}

func (b *BookHandler) RedeemBook(c *fiber.Ctx) error {
	var donatedBook entity.DonatedBook
	if err := c.BodyParser(&donatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	model := usecase.NewDonateBookModel(&donatedBook)
	if err := usecase.NewDonateBookUseCase(b.BookRepo, b.DonatedBookRepo, model).Handle(); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Redirect("/")
}
