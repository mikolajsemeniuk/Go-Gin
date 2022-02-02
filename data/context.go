package data

import (
	"fmt"

	"github.com/mikolajsemeniuk/Supreme-Go/configuration"
	"gorm.io/driver/postgres"
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

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configuration.Config.GetString("database.host"),
		configuration.Config.GetString("database.username"),
		configuration.Config.GetString("database.password"),
		configuration.Config.GetString("database.databasename"),
		configuration.Config.GetString("database.port"))

	var err error
	Context, err = gorm.Open(postgres.Open(connectionString), config)

	if err != nil {
		panic("failed to connect database")
	}
}
