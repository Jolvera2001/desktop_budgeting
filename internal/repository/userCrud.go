package repository

import (
	"desktop_budgeting/internal/models"

	"gorm.io/gorm"
)

type UserCrudInterface interface {
	Create(user *models.User) (uint, error)
	Get(id uint) (*models.User, error)
	GetMany() ([]models.User, error)
	Update(user *models.User) error // must contain Id for specific row to update
	Delete(id uint) error
}

type UserCrud struct {
	repo *gorm.DB
}

func (this *UserCrud) Create(user *models.User) (uint, error) {
	res := this.repo.Create(user)
	if res.Error != nil {
		return 0, res.Error
	}
	return user.ID, nil
}

func (this *UserCrud) Get(id uint) (*models.User, error) {
	var user models.User
	res := this.repo.First(&user, id)
	if res.Error != nil {
		return &models.User{}, res.Error
	}
	return &user, nil
}

func (this *UserCrud) GetMany() ([]models.User, error) {
	var users []models.User
	res := this.repo.First(&users)
	if res != nil {
		return []models.User{}, res.Error
	}
	return users, nil
}

func (this *UserCrud) Update(user *models.User) error {
	// assumes user contains ID
	res := this.repo.Save(user)
	if res != nil {
		return res.Error
	}
	return nil
}

func (this *UserCrud) Delete(id uint) error {
	res := this.repo.Delete(&models.User{}, id)
	if res != nil {
		return res.Error
	}
	return nil
}
