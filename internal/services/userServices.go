package services

import (
	m "desktop_budgeting/internal/models"
	r "desktop_budgeting/internal/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	crud r.UserCrudInterface
}

func (s *UserService) Register(dto m.UserDto) (*m.User, error) {
	hashedPass, err := HashPassword(dto.Password)
	if err != nil {
		return &m.User{}, err
	}

	newUser := m.User{
		Name: dto.Name,
		Email: dto.Email,
		Password: hashedPass,
	}

	_, err = s.crud.Create(&newUser)
	if err != nil {
		return &m.User{}, err
	}

	return &newUser, nil
}

func (s *UserService) Login(id uint, password, hash string) (*m.User, error) {
	var user *m.User
	if password == "" || hash == "" {
		return &m.User{}, fmt.Errorf("no values present")
	}

	check := CheckHash(password, hash)
	if !check {
		return &m.User{}, fmt.Errorf("password does not match")
	}

	user, err := s.crud.Get(id)
	if err != nil {
		return &m.User{}, err
	}
	return user, nil
}

func (s *UserService) GetAllProfiles() ([]*m.User, error) {
	users, err := s.crud.GetMany()
	if err != nil {
		return []*m.User{}, err
	}

	return users, nil
}

func (s *UserService) UpdateProfile(id uint, dto m.UserDto) error {
	updatedUser := m.User{}
}

func (s *UserService) DeleteProfile(id uint) error {
	err := s.crud.Delete(id)
	return err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
