package repository

import "desktop_budgeting/internal/models"

type CategoryCrudInterface interface {
	Create(category *models.Category) (uint, error)
	Get(id uint) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id uint) error
}