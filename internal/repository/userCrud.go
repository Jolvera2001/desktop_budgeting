package repository

import "desktop_budgeting/internal/models"

type UserCrudInterface interface {
	Create(user *models.User) uint
	Get(id uint) (*models.User, error)
	GetMany() []models.User
	Update(user *models.User) error // must contain Id for specific row to update
	Delete(id uint) error
}
