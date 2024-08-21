package book

import (
	"web/model"

	"gorm.io/gorm"
)

type Book struct {
	ID    int64
	Title string
}

type BookDAO struct {
	DB *gorm.DB
}

func NewBookDAO() BookDAO {
	return BookDAO{DB: model.DB2()}
}

func (Book) TableName() string {
	return "book"
}

func (dao *BookDAO) Insert(book Book) int64 {
	result := dao.DB.Create(&book)
	if result.Error != nil {
		panic(result.Error)
	}
	return book.ID
}

func (dao *BookDAO) GetByID(id int64) (Book, error) {
	var book Book
	result := dao.DB.Where("id = ?", id).First(&book)
	return book, result.Error
}

func (dao *BookDAO) GetByIDs(ids []int64) []Book {
	var books []Book
	result := dao.DB.Where("id in (?)", ids).Find(&books)
	if result.Error != nil {
		panic(result.Error)
	}
	return books
}
