package repository

import (
	"github.com/ArtuoS/doa-livros/internal/entity"
	"github.com/jmoiron/sqlx"
)

type BookRepository struct {
	DB *sqlx.DB
}

func (r *BookRepository) CreateBook(book *entity.Book) error {
	err := r.DB.QueryRowx(
		"INSERT INTO books (title, author, user_id, donating) VALUES ($1, $2, $3, $4) RETURNING id",
		book.Title, book.Author, book.UserId, book.Donating).Scan(&book.Id)
	return err
}

func (r *BookRepository) GetBook(id int64) (entity.Book, error) {
	var book entity.Book
	err := r.DB.Get(&book, "SELECT id, title, author, user_id, donating FROM books WHERE id=$1", id)
	return book, err
}

func (r *BookRepository) GetAllBooks() ([]entity.Book, error) {
	var books []entity.Book
	err := r.DB.Select(&books, "SELECT id, title, author, user_id, donating FROM books ORDER BY donating DESC")
	return books, err
}

func (r *BookRepository) UpdateBook(book *entity.Book) error {
	_, err := r.DB.Exec(
		"UPDATE books SET title=$1, author=$2 WHERE id=$3",
		book.Title, book.Author, book.Id)
	return err
}

func (r *BookRepository) DeleteBook(id int64) error {
	_, err := r.DB.Exec("DELETE FROM books WHERE id=$1", id)
	return err
}

func (r *BookRepository) ListBooksByUser(userId int64) ([]entity.Book, error) {
	var books []entity.Book
	err := r.DB.Select(&books, "SELECT id, title, author, donating FROM books WHERE user_id=$1", userId)
	return books, err
}

func (r *BookRepository) SetDonatingBook(donating bool, bookId int64) error {
	_, err := r.DB.Exec(
		"UPDATE books SET donating=$1 WHERE id=$2",
		donating, bookId)
	return err
}

func (r *BookRepository) ChangeOwner(userId, bookId int64) error {
	_, err := r.DB.Exec(
		"UPDATE books SET user_id=$1, donating=$2 WHERE id=$3",
		userId, false, bookId)

	return err
}
