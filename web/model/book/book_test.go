package book_test

import (
	"testing"

	"web/model"
	book_model "web/model/book"

	"github.com/stretchr/testify/assert"
)

func TestBookDAO(t *testing.T) {
	tx := model.DB2().Begin()
	book_dao := book_model.BookDAO{DB: tx}
	book_id := book_dao.Insert(book_model.Book{Title: "test"})
	book := book_dao.GetByID(book_id)
	assert.Equal(t, "test", book.Title)

	book_id2 := book_dao.Insert(book_model.Book{Title: "test2"})
	books := book_dao.GetByIDs([]int64{book_id, book_id2})
	assert.Equal(t, 2, len(books))
	assert.Equal(t, book_id, books[0].ID)
	assert.Equal(t, book_id2, books[1].ID)
	tx.Rollback()
}
