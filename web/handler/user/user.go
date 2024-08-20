package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
	data := fmt.Sprint("user:", c.Param("id"))
	c.String(200, data)
}