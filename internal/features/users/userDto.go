package users

type UserDto struct {
	Email        string           `json:"email"`
	Name         string           `json:"name"`
	BudgetPeriod budgetPeriodEnum `json:"budget_period"`
}
