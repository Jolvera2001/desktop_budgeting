package repository

import (
	"desktop_budgeting/internal/models"

	"gorm.io/gorm"
)

type IncomeCrudInterface interface {
	Create(income *models.Income) (uint, error)
	Get(id uint) (*models.Income, error)
	Update(income *models.Income) error
	Delete(id uint) error
}

type IncomeCrud struct {
	Repo *gorm.DB
}

func NewIncomeCrud(repo *gorm.DB) *IncomeCrud {
	return &IncomeCrud{Repo: repo}
}

func (c *IncomeCrud) Create(income *models.Income) (uint, error) {
	res := c.Repo.Create(income)
	if res.Error != nil {
		return 0, res.Error
	}
	return income.ID, nil
}

func (c *IncomeCrud) Get(id uint) (*models.Income, error) {
	var income models.Income
	res := c.Repo.First(&income, id)
	if res.Error != nil {
		return &models.Income{}, res.Error
	}
	return &income, nil
}

func (c *IncomeCrud) Update(income *models.Income) error {
	res := c.Repo.Save(income)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *IncomeCrud) Delete(id uint) error {
	res := c.Repo.Delete(&models.Income{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
