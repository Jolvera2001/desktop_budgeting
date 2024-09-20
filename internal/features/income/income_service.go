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

}

func (s *IncomeService) GetIncomeList(userId int64) ([]Income, error) {

}

func (s *IncomeService) UpdateIncome(update Income) error {

}

func (s *IncomeService) DeleteIncome(id int64) error {

}
