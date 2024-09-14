package budgets

type IBudgetService interface {
	CreateBudget(dto BudgetDto) (int64, error)
	GetBudget(id int64) (Budget, error)
	GetBudgets(userId int64) ([]Budget, error)
	UpdateBudget(update Budget) error
	DeleteBudget(id int64) error
}