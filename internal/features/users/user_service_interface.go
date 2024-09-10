package users

type IUserService interface {
	CreateUser(dto UserDto) (int64, error)
	GetUser(id int64) (User, error)
	GetUsers() ([]User, error)
	UpdateUser(update User) error
	DeleteUser(id int64) error
}