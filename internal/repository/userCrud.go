package repository

import (
	"desktop_budgeting/internal/models"

	"gorm.io/gorm"
)

type UserCrudInterface interface {
	Create(user *models.User) (uint, error)
	Get(id uint) (*models.User, error)
	GetMany() ([]*models.User, error)
	Update(user *models.User) error // must contain Id for specific row to update
	Delete(id uint) error
}

type UserCrud struct {
	Repo *gorm.DB
}

func NewUserCrud(repo *gorm.DB) *UserCrud {
	return &UserCrud{Repo: repo}
}

func (c *UserCrud) Create(user *models.User) (uint, error) {
	res := c.Repo.Create(user)
	if res.Error != nil {
		return 0, res.Error
	}
	return user.ID, nil
}

func (c *UserCrud) Get(id uint) (*models.User, error) {
	var user models.User
	res := c.Repo.Preload("Budgets").Preload("Budgets").First(&user, id)
	if res.Error != nil {
		return &models.User{}, res.Error
	}
	return &user, nil
}

func (c *UserCrud) GetMany() ([]*models.User, error) {
	var users []*models.User
	res := c.Repo.Find(&users)
	if res.Error != nil {
		return []*models.User{}, res.Error
	}
	return users, nil
}

func (c *UserCrud) Update(user *models.User) error {
	// assumes user contains ID
	res := c.Repo.Save(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *UserCrud) Delete(id uint) error {
	res := c.Repo.Delete(&models.User{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
