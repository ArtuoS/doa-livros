package repository

import (
	"github.com/ArtuoS/doa-livros/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	err := r.DB.QueryRowx(
		"INSERT INTO users (first_name, last_name) VALUES ($1, $2) RETURNING id",
		user.FirstName, user.LastName).Scan(&user.Id)
	return err
}

func (r *UserRepository) GetUser(id int64) (entity.User, error) {
	var user entity.User
	err := r.DB.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return user, err
	}

	err = r.DB.Select(&user.Books, "SELECT * FROM books WHERE user_id=$1", id)
	if err != nil {
		return user, err
	}

	err = r.DB.Select(&user.DonatedBooks, "SELECT * FROM donated_books WHERE from_user_id=$1", id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(user *entity.User) error {
	_, err := r.DB.Exec(
		"UPDATE users SET first_name=$1, last_name=$2 WHERE id=$3",
		user.FirstName, user.LastName, user.Id)
	return err
}

func (r *UserRepository) DeleteUser(id int64) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
