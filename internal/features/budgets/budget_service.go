package budgets

import "database/sql"

type BudgetService struct {
	Client *sql.DB
}

func (b *BudgetService) CreateBudget() (int64, error) {

}

func (b *BudgetService) GetBudget() (Budget, error) {

}

func (b *BudgetService) GetBudgets() ([]Budget, error) {

}

func (b *BudgetService) UpdateBudget() error {

}

func (b *BudgetService) DeleteBudget() error {

}
