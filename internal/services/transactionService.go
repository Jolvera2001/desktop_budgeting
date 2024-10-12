package services

import (
	m "desktop_budgeting/internal/models"
	r "desktop_budgeting/internal/repository"
) 

type TransactionService struct {
	Crud r.TransactionCrudInterface
}

func NewTransactionService(crud r.TransactionCrudInterface) *TransactionService {
	if crud == nil {
		panic("crud interface cannot be nil")
	}
	return &TransactionService{Crud: crud}
}

func (s *TransactionService) Add(dto m.TransactionDto) (*m.Transaction, error) {
	newTransaction := m.Transaction{
		BudgetID: dto.BudgetID,
		Description: dto.Description,
		Amount: dto.Amount,
		Date: dto.Date,
	}

	_, err := s.Crud.Create(&newTransaction)
	if err != nil {
		return &m.Transaction{}, err
	}
	return &newTransaction, nil
}

func (s *TransactionService) Check(id uint) (*m.Transaction, error) {
	transaction, err := s.Crud.Get(id)
	if err != nil {
		return &m.Transaction{}, err
	}
	return transaction, nil
}

func (s *TransactionService) Revise(id uint, dto m.TransactionDto) error {
	transaction, err := s.Crud.Get(id)
	if err != nil {
		return err
	}

	if dto.BudgetID != 0 {
		transaction.BudgetID = dto.BudgetID
	}
	if dto.Amount != 0 {
		transaction.Amount = dto.Amount
	}
	if dto.Date.IsZero() {
		transaction.Date = dto.Date
	}
	if dto.Description != "" {
		transaction.Description = dto.Description
	}

	if err := s.Crud.Update(transaction); err != nil {
		return err
	}

	return nil
}

func (s *TransactionService)  Remove(id uint) error {
	err := s.Crud.Delete(id)
	if err != nil {
		return err
	}
	return nil
}



