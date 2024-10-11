package repository

import (
	"desktop_budgeting/internal/models"

	"gorm.io/gorm"
)

type CategoryCrudInterface interface {
	Create(category *models.Category) (uint, error)
	Get(id uint) (*models.Category, error)
	Update(category *models.Category) error
	Delete(id uint) error
}

type CategoryCrud struct {
	Repo *gorm.DB
}

func NewCategoryCrud(repo *gorm.DB) *CategoryCrud {
	return &CategoryCrud{Repo: repo}
}

func (c *CategoryCrud) Create(category *models.Category) (uint, error) {
	res := c.Repo.Create(category)
	if res.Error != nil {
		return 0, res.Error
	}
	return category.ID, nil
}

func (c *CategoryCrud) Get(id uint) (*models.Category, error) {
	var category models.Category
	res := c.Repo.First(&category, id)
	if res.Error != nil {
		return &models.Category{}, res.Error
	}
	return &category, nil
}

func (c *CategoryCrud) Update(category *models.Category) error {
	res := c.Repo.Save(category)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *CategoryCrud) Delete(id uint) error {
	res := c.Repo.Delete(&models.Category{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

