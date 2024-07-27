package book

import (
	"github.com/ArtuoS/doa-livros/entity"
	"github.com/ArtuoS/doa-livros/infrastructure/repository"
)

type Service struct {
	bookRepo repository.BookRepository
}

func NewService(bookRepo repository.BookRepository) *Service {
	return &Service{
		bookRepo: bookRepo,
	}
}

func (s *Service) CreateBook(book *entity.Book) error {
	return s.bookRepo.CreateBook(book)
}

func (s *Service) GetBook(id int64) (entity.Book, error) {
	return s.bookRepo.GetBook(id)
}

func (s *Service) GetAllBooks() ([]entity.Book, error) {
	return s.bookRepo.GetAllBooks()
}

func (s *Service) UpdateBook(book *entity.Book) error {
	return s.bookRepo.UpdateBook(book)
}

func (s *Service) DeleteBook(id int64) error {
	return s.bookRepo.DeleteBook(id)
}

func (s *Service) ListBooksByUser(userId int64) ([]entity.Book, error) {
	return s.bookRepo.ListBooksByUser(userId)
}

func (s *Service) SetDonatingBook(donating bool, bookId int64) error {
	return s.bookRepo.SetDonatingBook(donating, bookId)
}

func (s *Service) ChangeOwner(userId, bookId int64) error {
	return s.bookRepo.ChangeOwner(userId, bookId)
}

func (s *Service) AddBookToDonation(bookId int64) error {
	return s.bookRepo.AddBookToDonation(bookId)
}
