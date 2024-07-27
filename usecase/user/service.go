package user

import (
	"github.com/ArtuoS/doa-livros/entity"
	"github.com/ArtuoS/doa-livros/infrastructure/repository"
)

type Service struct {
	repo repository.UserRepository
}

func NewService(userRepo repository.UserRepository) *Service {
	return &Service{
		repo: userRepo,
	}
}

func (s *Service) CreateUser(user *entity.User) error {
	return s.repo.CreateUser(user)
}

func (s *Service) GetUser(id int64) (entity.User, error) {
	return s.repo.GetUser(id)
}

func (s *Service) GetUserByAuth(auth entity.Auth) (entity.User, error) {
	return s.repo.GetUserByAuth(auth)
}

func (s *Service) UpdateUser(user *entity.User) error {
	return s.repo.UpdateUser(user)
}

func (s *Service) DeleteUser(id int64) error {
	return s.repo.DeleteUser(id)
}
