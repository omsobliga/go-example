package book_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"web"
	"web/handler"
	book_model "web/model/book"
)

func TestBookHandler(t *testing.T) {
	router := web.GetRouter()

	book_dao := book_model.NewBookDAO()
	book_id := book_dao.Insert(book_model.Book{Title: "test"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprint("/books/", book_id), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response handler.Book
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)
	assert.Equal(t, response.ID, int64(book_id))
}
