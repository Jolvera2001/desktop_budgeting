package repository

import "desktop_budgeting/internal/models"

type IncomeCrudInterface interface {
	Create(income *models.Income) (uint, error)
	Get(id uint) (*models.Income, error)
	Update(income *models.Income) error
	Delete(id uint) error
}