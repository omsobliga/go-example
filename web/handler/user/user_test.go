package user_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"web"
	"web/handler"
	user_model "web/model/user"
)

func TestUserHandler(t *testing.T) {
	router := web.GetRouter()

	user_dao := user_model.UserDAO{}
	user_id := user_dao.Insert("test", 10)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprint("/users/", user_id), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response handler.User
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, response.ID, int64(user_id))
}


func TestAsyncUserHandler(t *testing.T) {
	router := web.GetRouter()

	user_dao := user_model.UserDAO{}
	user_id := user_dao.Insert("test", 10)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprint("/users/", user_id), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response handler.User
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, response.ID, int64(user_id))

	// NotFound
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", fmt.Sprint("/users/", 13413143124), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response2 handler.ErrorMessage
	err = json.Unmarshal([]byte(w.Body.String()), &response2)
	assert.Nil(t, err)
	assert.Equal(t, response2.Error, "NotFound")
}
