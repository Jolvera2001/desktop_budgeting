package models

type budgetPeriodEnum int

const (
	Monthly = iota + 1
	Biweekly
	weekly
)

type User struct {
	ID           uint             `json:"_id" gorm:"primaryKey"`
	Email        string            `json:"email"`
	Name         string            `json:"name"`
	BudgetPeriod *budgetPeriodEnum `json:"budget_period"`
}
