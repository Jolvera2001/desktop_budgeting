package budgets

import (
	"database/sql"
	"fmt"
)

type BudgetService struct {
	Client *sql.DB
}

func (s *BudgetService) CreateBudget(dto BudgetDto) (int64, error) {
	res, err := s.Client.Exec("INSERT INTO budgets (userId, name, category, amount) VALUES (?, ?, ?, ?)",
		dto.UserID, dto.Name, dto.Category, dto.Amount)
	if err != nil {
		return 0, fmt.Errorf("add budget: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add budget: %v", err)
	}

	return id, nil
}

func (s *BudgetService) GetBudget(id int64) (Budget, error) {
	var budget Budget
	row := s.Client.QueryRow("SELECT * FROM budgets WHERE id = ?", id)

	if err := row.Scan(&budget.ID, &budget.UserID, &budget.Name, &budget.Category, &budget.Amount); err != nil {
		return Budget{}, fmt.Errorf("error fetching budget: %v", err)
	}

	return budget, nil
}

func (s *BudgetService) GetBudgets(userId int64) ([]Budget, error) {
	var budgets []Budget

	rows, err := s.Client.Query("SELECT * FROM budgets WHERE userId = ?", userId)
	if err != nil {
		return []Budget{}, fmt.Errorf("error fetching budgets: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var budget Budget

		if err := rows.Scan(&budget.ID, &budget.UserID, &budget.Name, &budget.Category, &budget.Amount); err != nil {
			return nil, fmt.Errorf("error fetching budgets: %v", err)
		}

		budgets = append(budgets, budget)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while fetching budgets: %v", err)
	}

	return budgets, nil
}

func (s *BudgetService) UpdateBudget(update Budget) error {
	_, err := s.Client.Exec("UPDATE budgets SET name = ?, category = ?, amount = ? WHERE id = ?;", 
	update.Name, update.Category, update.Amount, update.ID)
	if err != nil {
		return fmt.Errorf("error updating budget: %v", err)
	}

	return nil
}

func (s *BudgetService) DeleteBudget(id int64) error {
	_, err := s.Client.Exec("DELETE FROM budgets WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error deleting budget: %v", err)
	}

	return nil
}
