package income

import "database/sql"

type IncomeDto struct {
	UserID    int          `json:"user_id"`
	Category  string       `json:"category"`
	Amount    float64      `json:"amount"`
	IsRegular bool         `json:"is_regular"`
	Date      sql.NullTime `json:"date"`
}
