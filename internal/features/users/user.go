package users

import "time"

type budgetPeriod int

const (
	Monthly = iota + 1
	Biweekly
	weekly
)

type User struct {
	ID           int          `json:"_id"`
	Email        string       `json:"email"`
	Name         string       `json:"name"`
	BudgetPeriod budgetPeriod `json:"budget_period"`
	BudgetStart  time.Time    `json:"budget_start"`
}
