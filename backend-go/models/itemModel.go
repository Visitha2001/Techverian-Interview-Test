package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ID    uint    `gorm:"primarykey;autoIncrement" json:"id"`
	Name  string  `gorm:"not null" json:"name"`
	Price float64 `gorm:"not null" json:"price"`
}

func MigrateItems(db *gorm.DB) error {
	err := db.AutoMigrate(&Item{})
	if err != nil {
		return err
	}
	return nil
}
