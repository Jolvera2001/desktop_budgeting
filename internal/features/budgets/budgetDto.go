package budgets

type BudgetDto struct {
	UserID   int     `json:"user_id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}
