package crud

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/tniita/play-with-go/repo"
)

const dbName = "handson.db"

func GetItems() []repo.Items {
	items := []repo.Items{}
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
	result := db.Find(&items)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return items
}

func GetItemById(id string) repo.Items {
	item := repo.Items{}
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
	result := db.First(&item, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return item
}

func UpdateItem(items repo.Items) int64 {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
	result := db.Save(&items)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return result.RowsAffected
}

func DeleteItem(id string) int64 {
	item := repo.Items{}
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}
	result := db.Delete(&item, id)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return result.RowsAffected
}
