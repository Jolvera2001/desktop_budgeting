package users

import(
	"desktop_budgeting/internal/database"
)

type UserService struct {
	db *database.SqliteClient
}

func (s *UserService) CreateUser(dto UserDto) (User, error) {

}

func (s *UserService) GetUsers() ([]User, error) {

}

func (s *UserService) UpdateUser(update User) (User, error) {

}

func (s *UserService) DeleteUser(id int) error {

}