package users

import (
	"database/sql"
)

type budgetPeriodEnum int

const (
	Monthly = iota + 1
	Biweekly
	weekly
)

type User struct {
	ID           int               `json:"_id"`
	Email        string            `json:"email"`
	Name         string            `json:"name"`
	BudgetPeriod *budgetPeriodEnum `json:"budget_period"`
	BudgetStart  sql.NullTime      `json:"budget_start"`
}
