package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // import
)

// DB database
var DB *gorm.DB

// Open connection and migrates
func Open() *gorm.DB {
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("error connecting to Database")
	}
	return DB
}
