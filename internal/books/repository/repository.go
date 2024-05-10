package repository

import (
	"main/data/entity"

	"gorm.io/gorm"
)

type BookRepo struct {
	bookDB *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{bookDB: db}

}

func (br *BookRepo) AddBook(book *entity.Book) (err error) {
	err = br.bookDB.Create(book).Error
	return
}

func (br *BookRepo) AddExistingBook(book *entity.Book) (err error) {
	err = br.bookDB.Save(book).Error
	return
}

func (br *BookRepo) GetBookByTitle(title string) (book *entity.Book, found bool) {
	book = new(entity.Book)
	err := br.bookDB.Where("title = ?", title).First(book).Error
	if err != nil {
		return nil, false
	}
	return book, true

}