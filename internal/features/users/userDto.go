package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
) 

type UserDto struct {
	Email        string           `json:"email"`
	Name         string           `json:"name"`
	BudgetPeriod budgetPeriodEnum `json:"budget_period"`
	BudgetStart  sql.NullTime     `json:"budget_start"`
}

func (u *UserDto) UnmarshalJSON(data []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	u.Email = raw["email"].(string)
	u.Name = raw["name"].(string)
	
	if bp, ok := raw["budget_period"].(float64); ok {
		u.BudgetPeriod = budgetPeriodEnum(int(bp)) // Convert float64 to int and then to budgetPeriodEnum
	} else {
		return fmt.Errorf("invalid budget_period format")
	}

	if bs, ok := raw["budget_start"].(string); ok {
		t, err := time.Parse(time.RFC3339, bs)
		if err != nil {
			return err
		}
		u.BudgetStart = sql.NullTime{Time: t, Valid: true}
	} else {
		u.BudgetStart = sql.NullTime{Valid: false}
	}

	return nil
}
