package income

import (
	"database/sql"
	"fmt"
)

type IncomeService struct {
	Client *sql.DB
}

func (s *IncomeService) CreateIncome(dto IncomeDto) (int64, error) {
	res, err := s.Client.Exec("INSERT INTO income (userId, category, amount, isRegular, date) VALUES (?, ?, ?, ?, ?)",
		dto.UserID, dto.Category, dto.Amount, dto.IsRegular, dto.Date)
	if err != nil {
		return 0, fmt.Errorf("add income: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add income: %v", err)
	}

	return id, nil
}

func (s *IncomeService) GetIncome(id int64) (Income, error) {
	var income Income
	row := s.Client.QueryRow("SELECT * FROM income WHERE id = ?", id)

	if err := row.Scan(&income.ID, &income.UserID, &income.CategoryID, &income.Amount, &income.IsRegular, &income.Date); err != nil {
		return Income{}, fmt.Errorf("error fetching income: %v", err)
	}

	return income, nil
}

func (s *IncomeService) GetIncomeList(userId int64) ([]Income, error) {
	var incomeList []Income

	rows, err := s.Client.Query("SELECT * FROM income WHERE id = ?", userId)
	if err != nil {
		return []Income{}, fmt.Errorf("error fetching income list: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var income Income

		if err := rows.Scan(&income.ID, &income.UserID, &income.CategoryID, &income.Amount, &income.IsRegular, &income.Date); err != nil {
			return nil, fmt.Errorf("error fetching income: %v", err)
		}

		incomeList = append(incomeList, income)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error fetching incomeList: %v", err)
	}

	return incomeList, nil
}

func (s *IncomeService) UpdateIncome(update Income) error {
	query := `
	UPDATE FROM income
	SET category = ?, amount = ?, isRegular = ?, date = ?
	WHERE id = ?;`

	_, err := s.Client.Exec(query, update.CategoryID, update.Amount, update.IsRegular, update.Date, update.ID)
	if err != nil {
		return fmt.Errorf("error updating income: %v", err)
	}

	return nil
}

func (s *IncomeService) DeleteIncome(id int64) error {
	_, err := s.Client.Exec("DELETE FROM income WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error deleting income: %v", err)
	}

	return nil
}
