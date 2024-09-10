package users

import (
	"database/sql"
	db "desktop_budgeting/internal/database"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testDB db.SqliteClient

func TestImplementInterface(t *testing.T) {
	var _ IUserService = &UserService{}
}

func TestCreateUser(t *testing.T) {
	setup()
	dtoUser := UserDto{
		Name:         "JohnGoat",
		Email:        "JohnGoat@test.com",
		BudgetPeriod: 1,
		BudgetStart: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	service := UserService{client: testDB.Db}

	id, err := service.CreateUser(dtoUser)
	assert.Nil(t, err, "process should not bring up an error")
	userList, err := service.GetUsers()

	assert.Nil(t, err, "error should be nil")
	assert.Greater(t, id, int64(0), "there should be an id to lead back to row")
	assert.NotEmpty(t, userList, "list should not be empty")
	
	tearDown(id)
}

func TestGetUser(t *testing.T) {
	setup()
	userToDelete := UserDto{
		Name:         "DeleteMe!",
		Email:        "DeleteME!@delete.com",
		BudgetPeriod: 1,
		BudgetStart: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	service := UserService{client: testDB.Db}

	id, err := service.CreateUser(userToDelete)
	assert.Nil(t, err, "process should not bring up an error")
	actualUser, err := service.GetUser(id)

	assert.Nil(t, err, "process should not bring up an error")
	assert.NotEqual(t, User{}, actualUser, "user should not be empty")

	tearDown(id)
}

func TestGetUsers(t *testing.T) {
	setup()
	var users []User
	service := UserService{client: testDB.Db}

	users, err := service.GetUsers()

	assert.Nil(t, err, "there should not be an error")
	assert.NotEmpty(t, users, "user list should not be empty")

	tearDown(0)
}

func TestUpdateUser(t *testing.T) {
	setup()
	dtoUser := UserDto{
		Name:         "JohnGoat",
		Email:        "JohnGoat@test.com",
		BudgetPeriod: 1,
		BudgetStart: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	service := UserService{client: testDB.Db}

	id, err := service.CreateUser(dtoUser)
	assert.Nil(t, err, "process should not bring up an error")
	userToChange, err := service.GetUser(id)
	assert.Nil(t, err, "process should not bring up an error")
	update := userToChange
	update.Name = "updated"
	err = service.UpdateUser(update)
	updatedUser, err := service.GetUser(id)

	assert.Nil(t, err, "there should not be an error")
	assert.NotEqual(t, updatedUser, userToChange, "users should not be the same")

	tearDown(id)
}

func TestDeleteUser(t *testing.T) {
	setup()
	userToDelete := UserDto{
		Name:         "DeleteMe!",
		Email:        "DeleteME!@delete.com",
		BudgetPeriod: 1,
		BudgetStart: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	service := UserService{client: testDB.Db}

	id, err := service.CreateUser(userToDelete)
	assert.Nil(t, err, "process should not bring up an error")
	err = service.DeleteUser(id)
	assert.Nil(t, err, "process should not bring up an error")
	actualUser, err := service.GetUser(id)

	assert.NotNil(t, err, "there should be an error")
	assert.Equal(t, User{}, actualUser, "struct should be empty")

	tearDown(0)
}

func setup() {
	testDB.ConnectToDB()
	testDB.SetUpDB()
}

func tearDown(id int64) {
	if id != 0 {
		testDB.Db.Exec("DELETE FROM users WHERE id = ?", id)
	}

	testDB.Db.Close()
}
