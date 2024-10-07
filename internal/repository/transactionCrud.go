package repository

import (
	"desktop_budgeting/internal/models"

	"gorm.io/gorm"
)

type TransactionCrudInterface interface {
	Create(transaction *models.Transaction) (uint, error)
	Get(id uint) (*models.Transaction, error)
	Update(transaction *models.Transaction) error
	Delete(id uint) error
}

type TransactionCrud struct {
	repo *gorm.DB
}

func (c *TransactionCrud) Create(transaction *models.Transaction) (uint, error) {
	res := c.repo.Create(transaction)
	if res.Error != nil {
		return 0, res.Error
	}
	return transaction.ID, nil
}

func (c *TransactionCrud) Get(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	res := c.repo.First(&transaction, id)
	if res.Error != nil {
		return &models.Transaction{}, res.Error
	}
	return &transaction, nil
}

func (c *TransactionCrud) Update(transaction *models.Transaction) error {
	res := c.repo.Save(transaction)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *TransactionCrud) Delete(id uint) error {
	res := c.repo.Delete(&models.Transaction{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

