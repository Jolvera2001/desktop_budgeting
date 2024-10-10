package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model `json:",inline"`
	Name       string `json:"category"`
}
