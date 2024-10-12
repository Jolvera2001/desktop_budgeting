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
	Repo *gorm.DB
}

func NewTransactionCrud(repo *gorm.DB) *TransactionCrud {
	return &TransactionCrud{Repo: repo}
}

func (c *TransactionCrud) Create(transaction *models.Transaction) (uint, error) {
	res := c.Repo.Create(transaction)
	if res.Error != nil {
		return 0, res.Error
	}
	return transaction.ID, nil
}

func (c *TransactionCrud) Get(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	res := c.Repo.First(&transaction, id)
	if res.Error != nil {
		return &models.Transaction{}, res.Error
	}
	return &transaction, nil
}

func (c *TransactionCrud) Update(transaction *models.Transaction) error {
	res := c.Repo.Save(transaction)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *TransactionCrud) Delete(id uint) error {
	res := c.Repo.Delete(&models.Transaction{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

