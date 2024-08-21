package user

import (
	"strconv"

	"web/handler"
	"web/model/user"

	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, handler.ErrorMessage{Error: "InvalidParam", Message: err.Error()})
		return
	}
	dao := user.NewUserDAO()
	user, err := dao.GetByID(int64(userID))
	if err != nil {
		c.JSON(200, handler.ErrorMessage{Error: "NotFound", Message: err.Error()})
		return
	} else {
		c.JSON(200, handler.User{ID: user.ID, Name: user.Name, Age: user.Age})
		return
	}
}

func AsyncUserHandler(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, handler.ErrorMessage{Error: "InvalidParam", Message: err.Error()})
		return
	}

	dao := user.NewUserDAO()
	chUser := make(chan *user.User)
	go dao.AsyncGetByID(int64(userID), chUser)

	user := <-chUser
	if user != nil {
		c.JSON(200, handler.User{ID: user.ID, Name: user.Name, Age: user.Age})
		return
	} else {
		c.JSON(200, handler.ErrorMessage{Error: "NotFound", Message: err.Error()})
		return
	}
}
