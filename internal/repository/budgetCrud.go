package repository

import (
	"desktop_budgeting/internal/models"

	"gorm.io/gorm"
)

type BudgetCrudInterface interface {
	Create(budget *models.Budget) (uint, error)
	Get(id uint) (*models.Budget, error)
	Update(budget *models.Budget) error
	Delete(id uint) error
}

type BudgetCrud struct { 
	Repo *gorm.DB
}

func NewBudgetCrud(repo *gorm.DB) *BudgetCrud {
	return &BudgetCrud{Repo: repo}
}

func (c *BudgetCrud) Create(budget *models.Budget) (uint, error) {
	res := c.Repo.Create(budget)
	if res.Error != nil {
		return 0, res.Error
	}
	return budget.ID, nil
}

func (c *BudgetCrud) Get(id uint) (*models.Budget, error) {
	var budget models.Budget
	res := c.Repo.Preload("Transactions").First(&budget, id)
	if res.Error != nil {
		return &models.Budget{}, res.Error
	}
	return &budget, nil
}

func (c *BudgetCrud) Update(budget *models.Budget) error {
	res := c.Repo.Save(budget)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *BudgetCrud) Delete(id uint) error {
	res := c.Repo.Delete(&models.Budget{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}


