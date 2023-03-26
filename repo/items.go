package repo

import (
	"gorm.io/gorm"
)

type Items struct {
	gorm.Model
	Name     string
	Desc     string
	Date     string
	PostedBy string
	BoughtBy string
	Price    string
	Status   string
}
