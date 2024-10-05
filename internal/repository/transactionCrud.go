package repository

type TransactionCrudInterface interface {
	Create()
	Get()
	Update()
	Delete()
}