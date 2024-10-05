package income

import "database/sql"

type IncomeDto struct {
	UserID     int64        `json:"user_id"`
	CategoryID int64        `json:"category"`
	Amount     float64      `json:"amount"`
	IsRegular  bool         `json:"is_regular"`
	Date       sql.NullTime `json:"date"`
}
