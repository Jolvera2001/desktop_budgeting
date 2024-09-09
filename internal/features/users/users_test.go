package users

import (
	"desktop_budgeting/internal/database"
	"testing"
)

func TestImplementInterface(t *testing.T) {
	var _ IUserService = &UserService{client: &database.SqliteClient{}}
}

func TestCreateUser(t *testing.T) {

}

func TestGetUsers(t *testing.T) {
	
}

func TestUpdateUser(t *testing.T) {
	
}

func TestDeleteUser(t *testing.T) {
	
}
// create user
// edit user
// delete user
// get users
