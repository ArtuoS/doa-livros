package repository

import (
	"github.com/ArtuoS/doa-livros/internal/entity"
	"github.com/jmoiron/sqlx"
)

type DonatedBookRepository struct {
	DB *sqlx.DB
}

func (r *DonatedBookRepository) CreateDonatedBook(donatedBook *entity.DonatedBook) error {
	err := r.DB.QueryRowx(
		"INSERT INTO donated_books (from_user_id, to_user_id, book_id) VALUES ($1, $2, $3)",
		donatedBook.FromUserId, donatedBook.ToUserId, donatedBook.BookId).Err()

	return err
}
