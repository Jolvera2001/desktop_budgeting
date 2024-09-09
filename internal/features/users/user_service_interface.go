package users

type IUserService interface {
	CreateUser(dto UserDto) (User, error)
	GetUsers() ([]User, error)
	UpdateUser(update User) (User, error)
	DeleteUser(id int) error
}