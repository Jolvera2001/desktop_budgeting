package services

import (
	m "desktop_budgeting/internal/models"
	r "desktop_budgeting/internal/repository"
)

type CategoryService struct {
	Crud r.CategoryCrudInterface
}

func NewCategoryService(crud r.CategoryCrudInterface) *CategoryService {
	if crud == nil {
		panic("crud interface cannot be nil")
	}
	return &CategoryService{Crud: crud}
}

func (s *CategoryService) Add(name string) (*m.Category, error) {
	newCategory := m.Category{
		Name: name,
	}

	_, err := s.Crud.Create(&newCategory)
	if err != nil {
		return &m.Category{}, err
	}
	return &newCategory, nil
}

func (s *CategoryService) Get(id uint) (*m.Category, error) {
	category, err := s.Crud.Get(id)
	if err != nil {
		return &m.Category{}, err
	}

	return category, nil
}

func (s *CategoryService) Rename(id uint, name string) error {
	category, err := s.Crud.Get(id)
	if err != nil {
		return err
	}

	if name != "" {
		category.Name = name
	}

	if err := s.Crud.Update(category); err != nil {
		return err
	}

	return nil


}

func (s *CategoryService) Delete(id uint) error {
	if err := s.Crud.Delete(id); err != nil {
		return err
	}
	return nil
}

