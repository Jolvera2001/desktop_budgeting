package budgets

type Budget struct {
	ID       int64   `json:"_id"`
	UserID   int64   `json:"user_id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}
