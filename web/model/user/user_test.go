package user_test

import (
	"testing"
	user_model "web/model/user"

	"github.com/stretchr/testify/assert"
)

func TestUserDAO(t *testing.T) {
	user_dao := user_model.UserDAO{}
	user_id := user_dao.Insert("test", 10)

	user := user_dao.GetByID(user_id)
	assert.Equal(t, "test", user.Name)
	assert.Equal(t, int32(10), user.Age)

	user_id2 := user_dao.Insert("test2", 10)
	users := user_dao.GetByIDs([]int64{user_id, user_id2})
	assert.Equal(t, 2, len(users))
	assert.Equal(t, user_id, users[0].ID)
	assert.Equal(t, user_id2, users[1].ID)
}
