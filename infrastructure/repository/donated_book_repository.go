package repository

import (
	"github.com/ArtuoS/doa-livros/entity"
	"github.com/jmoiron/sqlx"
)

type DonatedBookRepository struct {
	DB *sqlx.DB
}

func (r *DonatedBookRepository) CreateDonatedBook(donatedBook *entity.DonatedBook) error {
	err := r.DB.QueryRowx(
		"INSERT INTO donated_books (from_user_id, to_user_id, book_id, to_user_name) VALUES ($1, $2, $3, $4)",
		donatedBook.FromUserId, donatedBook.ToUserId, donatedBook.BookId, donatedBook.ToUserName).Err()

	return err
}
