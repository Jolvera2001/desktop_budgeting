package budgets

type BudgetDto struct {
	UserID     int64   `json:"user_id"`
	CategoryID int64   `json:"category"`
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
}
