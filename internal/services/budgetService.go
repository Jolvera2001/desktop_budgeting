package services

import (
	m "desktop_budgeting/internal/models"
	r "desktop_budgeting/internal/repository"
	"fmt"
)

type BudgetService struct {
	Crud r.BudgetCrudInterface
}

func NewBudgetService(crud r.BudgetCrudInterface) *BudgetService {
	if crud == nil {
		panic("crud interface cannot be nil")
	}
	return &BudgetService{Crud: crud}
}

func (s *BudgetService) Make(dto m.BudgetDto) (*m.Budget, error) {
	newBudget := m.Budget{
		UserID: dto.UserID,
		Name: dto.Name,
		Amount: dto.Amount,
	}

	_, err := s.Crud.Create(&newBudget)
	if err != nil {
		return &m.Budget{}, err
	}
	return &newBudget, nil
}

func (s *BudgetService) Check(id uint) (*m.Budget, error) {
	budget, err := s.Crud.Get(id)
	if err != nil {
		return &m.Budget{}, err
	}
	return budget, nil
}

func (s *BudgetService) Revise(id uint, dto m.BudgetDto) error {
	budget, err := s.Crud.Get(id)
	if err != nil {
		return fmt.Errorf("failed to fetch Budget: %w", err)
	}

	if dto.Name != "" {
		budget.Name = dto.Name
	}
	if dto.Amount != 0 {
		budget.Amount = dto.Amount
	}

	if err := s.Crud.Update(budget); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (s *BudgetService) Remove(id uint) error {
	err := s.Crud.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

