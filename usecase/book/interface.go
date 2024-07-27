package book

import "github.com/ArtuoS/doa-livros/entity"

type UseCase interface {
	CreateBook(book *entity.Book) error
	GetBook(id int64) (entity.Book, error)
	GetAllBooks() ([]entity.Book, error)
	UpdateBook(book *entity.Book) error
	DeleteBook(id int64) error
	ListBooksByUser(userId int64) ([]entity.Book, error)
	SetDonatingBook(donating bool, bookId int64) error
	ChangeOwner(userId, bookId int64) error
	AddBookToDonation(bookId int64) error
}
