package repository

import "desktop_budgeting/internal/models"

type BudgetCrudInterface interface {
	Create(budget *models.Budget) (uint, error)
	Get(id uint) (*models.Budget, error)
	Update(budget *models.Budget) error
	Delete(id uint) error
}