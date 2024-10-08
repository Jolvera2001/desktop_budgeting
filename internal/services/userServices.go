package services

import (
	m "desktop_budgeting/internal/models"
	r "desktop_budgeting/internal/repository"
)

type UserService struct {
	crud r.UserCrudInterface
}

func (s *UserService) CreateProfile() (uint, error) {

}

func (s *UserService) GetProfile(id uint) (*m.User, error) {

}

func (s *UserService) GetAllProfiles() {

}

func (s *UserService) UpdateProfile() {

}

func (s *UserService) DeleteProfile() {

}
