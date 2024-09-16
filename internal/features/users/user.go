package users

type budgetPeriodEnum int

const (
	Monthly = iota + 1
	Biweekly
	weekly
)

type User struct {
	ID           int64             `json:"_id"`
	Email        string            `json:"email"`
	Name         string            `json:"name"`
	BudgetPeriod *budgetPeriodEnum `json:"budget_period"`
}
