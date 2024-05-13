package usecase

import (
	"main/data/entity"
	"main/internal/books/repository"

	"github.com/google/uuid"
)

type BookUsecase struct {
	repo *repository.BookRepo
}

func NewBookUsecase(repo *repository.BookRepo) *BookUsecase {
	return &BookUsecase{repo: repo}
}

func (bu *BookUsecase) AddBook(book *entity.Book) (err error) {
	existing, found := bu.repo.GetBookByTitle(book.Title)
	if found {
		book.ID = existing.ID
		book.Stock += existing.Stock
		return bu.repo.AddExistingBook(book)
	}

	book.ID = uuid.New()
	return bu.repo.AddBook(book)
}

func (bu *BookUsecase) BrowseBook(request string) (books []entity.Book, err error) {
	return bu.repo.BrowseBook(request)
}