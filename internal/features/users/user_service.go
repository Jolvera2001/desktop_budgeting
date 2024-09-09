package users

import (
	"desktop_budgeting/internal/database"
	"fmt"
)

type UserService struct {
	client *database.SqliteClient
}

func (s *UserService) CreateUser(dto UserDto) (int64, error) {
	res, err := s.client.Db.Exec("INSERT INTO users (email, name) VALUES (?, ?)", dto.Email, dto.Name)
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
	row := s.client.Db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.BudgetStart, &user.BudgetPeriod); err != nil {
		return User{}, fmt.Errorf("error fetching user: %v", err)
	}

	return user, nil
}

func (s *UserService) GetUsers() ([]User, error) {
	var users []User

	rows, err := s.client.Db.Query("SELECT * FROM users")
	if err != nil {
		return []User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.BudgetStart, &user.BudgetPeriod); err != nil {
			return nil, fmt.Errorf("error fetching users: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) UpdateUser(update User) error {
	_, err := s.client.Db.Exec("UPDATE users SET name = ?, email = ?, budgetPeriod = ?, budgetStart = ? WHERE id = ?;",
		update.Name, update.Email, update.BudgetPeriod, update.BudgetStart, update.ID)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	return nil
}

func (s *UserService) DeleteUser(id int) error {
	_, err := s.client.Db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	return nil
}
