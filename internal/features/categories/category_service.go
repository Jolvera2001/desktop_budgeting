package categories

import "database/sql"

type CategoryService struct {
	Client *sql.DB
}

func (c *CategoryService) CreateCategory(dto string) (int64, error) {

}

func (c *CategoryService) GetCategory(id int64) (Category, error) {

}

func (c *CategoryService) GetCategoryList(userId int64) ([]Category, error) {

}

func (c *CategoryService) UpdateCategory(update Category) error {

}

func (c *CategoryService) DeleteCategory(id int64) error {

}
