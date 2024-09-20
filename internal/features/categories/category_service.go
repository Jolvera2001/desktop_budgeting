package categories

import (
	"database/sql"
	"fmt"
)

type CategoryService struct {
	Client *sql.DB
}

func (s *CategoryService) CreateCategory(dto CategoryDto) (int64, error) {
	res, err := s.Client.Exec("INSERT INTO categories (userId, name) VALUES (?, ?)",
		dto.UserID, dto.Name)
	if err != nil {
		return 0, fmt.Errorf("error adding category: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error adding category: %v", err)
	}

	return id, nil
}

func (s *CategoryService) GetCategory(id int64) (Category, error) {
	var category Category
	row := s.Client.QueryRow("SELECT * FROM categories WHERE id = ?", id)

	if err := row.Scan(&category.ID, &category.UserID, &category.Name); err != nil {
		return Category{}, fmt.Errorf("error getting category: %v", err)
	}

	return category, nil
}

func (s *CategoryService) GetCategoryList(userId int64) ([]Category, error) {
	var categories []Category

	rows, err := s.Client.Query("SELECT * FROM categories WHERE userId = ?", userId)
	if err != nil {
		return []Category{}, fmt.Errorf("error fetching categories: %v", err)

	}

	defer rows.Close()

	for rows.Next() {
		var category Category

		if err := rows.Scan(&category.ID, &category.UserID, category.Name); err != nil {
			return nil, fmt.Errorf("error fetching categories: %v", err)
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while fetching categories: %v", err)
	}

	return categories, nil
}

func (s *CategoryService) UpdateCategory(update Category) error {
	query := `
	UPDATE categories
	SET name = ?
	WHERE id = ?
	`
	_, err := s.Client.Exec(query, update.Name, update.ID)
	if err != nil {
		return fmt.Errorf("error updating category: %v", err)
	}

	return nil
}

func (s *CategoryService) DeleteCategory(id int64) error {
	_, err := s.Client.Exec("DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("error deleting category")
	}

	return nil
}
