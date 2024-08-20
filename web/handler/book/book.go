package book

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func BookHandler(c *gin.Context) {
	data := fmt.Sprint("book:", c.Param("id"))
	c.String(200, data)
}