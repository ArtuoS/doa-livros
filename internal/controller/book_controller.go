package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ArtuoS/doa-livros/internal/entity"
	"github.com/ArtuoS/doa-livros/internal/repository"
	"github.com/ArtuoS/doa-livros/internal/usecase"
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

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.BookRepo.CreateBook(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (c *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, err := c.BookRepo.GetBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (c *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
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
	if err := c.BookRepo.UpdateBook(&book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (c *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	if err := c.BookRepo.DeleteBook(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *BookController) RedeemBook(w http.ResponseWriter, r *http.Request) {
	var donatedBook entity.DonatedBook
	err := json.NewDecoder(r.Body).Decode(&donatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model := usecase.NewDonateBookModel(&donatedBook)
	if err := usecase.NewDonateBookUseCase(c.BookRepo, c.DonatedBookRepo, model).Handle(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
