package users

import (
	db "desktop_budgeting/internal/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDB db.SqliteClient

func TestImplementInterface(t *testing.T) {
	var _ IUserService = &UserService{}
}

func TestCreateUser(t *testing.T) {
	Setup()
	dtoUser := UserDto{
		Name:  "JohnGoat",
		Email: "JohnGoat@test.com",
	}
	service := UserService{client: testDB.Db}

	id, err := service.CreateUser(dtoUser)
	assert.Nil(t, err)
	userList, err := service.GetUsers()

	assert.Nil(t, err)
	assert.Greater(t, id, int64(0))
	assert.Greater(t, len(userList), 0)
	TearDown()
}

func TestGetUsers(t *testing.T) {
	Setup()
	TearDown()
}

func TestUpdateUser(t *testing.T) {
	Setup()
	TearDown()
}

func TestDeleteUser(t *testing.T) {
	Setup()
	TearDown()
}

func Setup() {
	testDB.ConnectToDB()
	testDB.SetUpDB()
}

func TearDown() {
	testDB.Db.Close()
}
