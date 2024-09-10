package users

import "database/sql"

type UserDto struct {
	Email        string            `json:"email"`
	Name         string            `json:"name"`
	BudgetPeriod budgetPeriodEnum `json:"budget_period"`
	BudgetStart  sql.NullTime      `json:"budget_start"`
}
