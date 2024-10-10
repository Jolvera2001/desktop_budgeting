package services

import (
	m "desktop_budgeting/internal/models"
	r "desktop_budgeting/internal/repository"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	crud r.UserCrudInterface
}

func (s *UserService) Register(dto m.UserDto) (*m.User, error) {
	hashedPass, err := hashPassword(dto.Password)
	if err != nil {
		return &m.User{}, err
	}

	newUser := m.User{
		Name:     dto.Name,
		Email:    dto.Email,
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

	check := checkHash(password, hash)
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
	user, err := s.crud.Get(id)
	if err != nil {
		return fmt.Errorf("failed to fetch user: %w", err)
	}

	if dto.Name != "" {
		user.Name = dto.Name
	}
	if dto.Email != "" {
		user.Email = dto.Email
	}
	if dto.Password != "" {
		hashedPass, err := hashPassword(dto.Password)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		user.Password = hashedPass
	}

	// Perform basic validation
	if err := validateUserFields(*user); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Update the user
	if err := s.crud.Update(user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (s *UserService) DeleteProfile(id uint) error {
	err := s.crud.Delete(id)
	return err
}

//  HELPER METHODS **********************************

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func validateUserFields(user m.User) error {
	if user.Name == "" {
		return errors.New("name cannot be empty")
	}
	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
	}
	// Add more validation as needed
	return nil
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@")
}
