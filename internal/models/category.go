package models

type Category struct {
	Base
	UserID    uint      `json:"userId"`
	Name      string    `json:"category"`
}
