package categories

type ICategoryService interface {
	CreatCategory(dto string) (int64, error)
	GeCategory(id int64) (Category, error)
	GetCategoryList(userId int64) ([]Category, error)
	UpdatCategory(update Category) error
	DeletCategory(id int64) error
}