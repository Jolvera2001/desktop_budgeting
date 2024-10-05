package repository

import "desktop_budgeting/internal/models"

type TransactionCrudInterface interface {
	Create(transaction *models.Transaction) uint
	Get(id uint) (*models.Transaction, error)
	Update(transaction *models.Transaction) error
	Delete(id uint) error
}