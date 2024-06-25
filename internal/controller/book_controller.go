package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ArtuoS/doa-livros/internal/entity"
	"github.com/ArtuoS/doa-livros/internal/repository"
	"github.com/ArtuoS/doa-livros/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
)

type BookController struct {
	BookRepo        *repository.BookRepository
	DonatedBookRepo *repository.DonatedBookRepository
}

func NewBookController(bookRepo *repository.BookRepository, donatedBookRepo *repository.DonatedBookRepository) *BookController {
	return &BookController{
		BookRepo:        bookRepo,
		DonatedBookRepo: donatedBookRepo,
	}
}

func (b *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := b.BookRepo.CreateBook(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (b *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, err := b.BookRepo.GetBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (b *BookController) GetAllBooks(c *fiber.Ctx) error {
	books, err := b.BookRepo.GetAllBooks()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Render("books", fiber.Map{"Books": books})
}

func (b *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book.Id = id
	if err := b.BookRepo.UpdateBook(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (b *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	if err := b.BookRepo.DeleteBook(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (b *BookController) RedeemBook(c *fiber.Ctx) error {
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
