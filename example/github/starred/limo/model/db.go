package model

import (
	"github.com/jinzhu/gorm"
	// Use the sqlite dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// InitDB initializes the database at the specified path
func InitDB(filepath string, verbose bool) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	db.LogMode(verbose)
	db.AutoMigrate(&Service{}, &Star{}, &Tag{})

	return db, nil
}
