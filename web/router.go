package web

import (
	"github.com/gin-gonic/gin"

	"web/handler/book"
	"web/handler/user"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users/:id", user.UserHandler)
	router.GET("/async/users/:id", user.AsyncUserHandler)
	router.GET("/books/:id", book.BookHandler)
	return router
}
