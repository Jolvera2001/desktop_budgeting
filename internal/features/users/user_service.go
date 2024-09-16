package users

import (
	"database/sql"
	"fmt"
)

type UserService struct {
	Client *sql.DB
}

func (s *UserService) CreateUser(dto UserDto) (int64, error) {
	res, err := s.Client.Exec("INSERT INTO users (email, name, budgetPeriod) VALUES (?, ?, ?)",
		dto.Email, dto.Name, dto.BudgetPeriod)
	if err != nil {
		return 0, fmt.Errorf("add user: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add user: %v", err)
	}

	return id, nil
}

func (s *UserService) GetUser(id int64) (User, error) {
	var user User
	row := s.Client.QueryRow("SELECT * FROM users WHERE id = ?", id)

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.BudgetPeriod); err != nil {
		return User{}, fmt.Errorf("error fetching user: %v", err)
	}

	return user, nil
}

func (s *UserService) GetUsers() ([]User, error) {
	var users []User

	rows, err := s.Client.Query("SELECT * FROM users")
	if err != nil {
		return []User{}, fmt.Errorf("error fetching users: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		var budgetPeriod sql.NullInt64

		if err := rows.Scan(&user.ID, &user.Email, &user.Name, &budgetPeriod); err != nil {
			return nil, fmt.Errorf("error fetching users: %v", err)
		}

		bp := budgetPeriodEnum(budgetPeriod.Int64)
		user.BudgetPeriod = &bp
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) UpdateUser(update User) error {
	_, err := s.Client.Exec("UPDATE users SET name = ?, email = ?, budgetPeriod = ? WHERE id = ?;",
		update.Name, update.Email, update.BudgetPeriod, update.ID)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	return nil
}

func (s *UserService) DeleteUser(id int64) error {
	_, err := s.Client.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	return nil
}
