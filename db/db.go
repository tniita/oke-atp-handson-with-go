package db

import (
	"time"

	"github.com/tniita/play-with-go/repo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() {
	db, err := gorm.Open(sqlite.Open("handson.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}

	db.AutoMigrate(&repo.Items{})

	db.Create(&repo.Items{
		Model:    gorm.Model{ID: 123},
		Name:     "Italian Antique Hand-Painted Porcelain Trinket Boxes",
		Desc:     "Antique hand-painted porcelain trinket boxes with hinged lids, purchased in 1985 at the Clignancourt Flea Market in Paris, mint condition, $75 each",
		Date:     time.Date(2018, 12, 20, 12, 34, 56, 123456, time.UTC).Format("20060102150405"),
		PostedBy: "6",
		BoughtBy: "",
		Price:    "70",
		Status:   "cancelled",
	})
}
