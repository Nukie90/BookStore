package usecase

import (
	"errors"
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

func (bu *BookUsecase) Checkout(userID uuid.UUID) (err error){
	var shoppingCart []entity.ShoppingCart
	shoppingCart, err = bu.repo.GetUserCart(userID)
	if err != nil {
		return
	}

	bookList := make([]entity.Book, len(shoppingCart))
	cost := 0.0

	for i, item := range shoppingCart {
		book := new(entity.Book)
		book, found := bu.repo.GetBookByID(item.BookID)
		if !found {
			return errors.New("book not found")
		}
		book.Stock = book.Stock - item.Quantity
		if book.Stock < 0 {
			return errors.New("not enough stock")
		}

		err = bu.repo.AddExistingBook(book)
		if err != nil {
			return err
		}

		bookList[i] = *book
		cost += item.Cost
	}
	
	err = bu.repo.AddSoldRecord(userID, bookList, cost)
	if err != nil {
		return
	}

	err = bu.repo.ClearUserCart(userID)
	if err != nil {
		return
	}

	return
}