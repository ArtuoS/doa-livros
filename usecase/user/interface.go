package user

import "github.com/ArtuoS/doa-livros/entity"

type UseCase interface {
	CreateUser(user *entity.User) error
	GetUser(id int64) (entity.User, error)
	GetUserByAuth(auth entity.Auth) (entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id int64) error
}
