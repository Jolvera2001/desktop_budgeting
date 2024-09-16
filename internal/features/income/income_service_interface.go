package income

type IIncomeService interface{
	CreateIncome(dto IncomeDto) (int64, error)
	GetIncome(id int64) (Income, error)
	GetIncomeList(userId int64) ([]Income, error)
	UpdateIncome(update Income) error
	DeleteIncome(id int64) error
}