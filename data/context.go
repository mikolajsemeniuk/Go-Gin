package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Context *gorm.DB
)

func init() {
	config := &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	}
	var err error
	Context, err = gorm.Open(sqlite.Open("test.db"), config)
	if err != nil {
		panic("failed to connect database")
	}
}
