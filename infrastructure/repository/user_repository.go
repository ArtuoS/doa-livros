package repository

import (
	"github.com/ArtuoS/doa-livros/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	err := r.DB.QueryRowx(
		"INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id",
		user.FirstName, user.LastName, user.Email, user.Password).Scan(&user.Id)
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

	err = r.DB.Select(&user.DonatedBooks, "SELECT db.*, CONCAT(u.first_name, ' ', u.last_name) as to_user_name FROM donated_books db INNER JOIN users u ON u.id = db.to_user_id WHERE from_user_id=$1", id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByAuth(auth entity.Auth) (*entity.User, error) {
	var user entity.User
	err := r.DB.Get(&user, "SELECT * FROM users WHERE email=$1 AND password=$2", auth.Email, auth.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
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
