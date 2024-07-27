package usecase

import (
	"github.com/ArtuoS/doa-livros/entity"
	"github.com/ArtuoS/doa-livros/infrastructure/repository"
)

type DonateBookUseCase struct {
	BookRepository        *repository.BookRepository
	DonatedBookRepository *repository.DonatedBookRepository
	Model                 *DonateBookModel
}

type DonateBookModel struct {
	DonateBook *entity.DonatedBook
}

func NewDonateBookUseCase(bookRepo *repository.BookRepository, donatedBookRepo *repository.DonatedBookRepository, model *DonateBookModel) *DonateBookUseCase {
	return &DonateBookUseCase{
		BookRepository:        bookRepo,
		DonatedBookRepository: donatedBookRepo,
		Model:                 model,
	}
}

func NewDonateBookModel(donatedBook *entity.DonatedBook) *DonateBookModel {
	return &DonateBookModel{
		DonateBook: donatedBook,
	}
}

func (d *DonateBookUseCase) Handle() error {
	err := d.BookRepository.ChangeOwner(d.Model.DonateBook.ToUserId, d.Model.DonateBook.BookId)
	if err != nil {
		return err
	}

	return d.DonatedBookRepository.CreateDonatedBook(d.Model.DonateBook)
}
