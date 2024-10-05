package repository

type UserCrudInterface interface {
	Create
	Get
	GetMany
	Update
	Delete
}
