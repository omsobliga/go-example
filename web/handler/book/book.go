package book

import (
	"strconv"

	"web/handler"
	"web/model/book"

	"github.com/gin-gonic/gin"
)

func BookHandler(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, handler.ErrorMessage{Error: "InvalidParam", Message: err.Error()})
		return
	}
	dao := book.NewBookDAO()
	book, err := dao.GetByID(int64(bookID))
	if err != nil {
		c.JSON(200, handler.ErrorMessage{Error: "NotFound", Message: err.Error()})
		return
	} else {
		c.JSON(200, handler.Book{ID: book.ID, Title: book.Title})
		return
	}

}
