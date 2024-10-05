package models

type Category struct {
	ID     int64  `json:"_id_"`
	UserID int64  `json:"userId"`
	Name   string `json:"category"`
}
