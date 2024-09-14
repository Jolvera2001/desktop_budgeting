package budgets

type BudgetDto struct {
	UserID   int     `json:"user_id"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}
