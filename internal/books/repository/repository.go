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

func (br *BookRepo) BrowseBook(request string) (books []entity.Book, err error) {
	switch request {
	case "all":
		err = br.bookDB.Find(&books).Error
	case "available":
		err = br.bookDB.Where("stock > 0").Find(&books).Error
	case "by_price":
		err = br.bookDB.Order("price desc").Find(&books).Error
	case "by_author":
		err = br.bookDB.Order("author").Find(&books).Error
	case "by_title":
		err = br.bookDB.Order("title").Find(&books).Error
	}
	return
}