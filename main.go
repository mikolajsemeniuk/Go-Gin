package main

import (
	"github.com/mikolajsemeniuk/Supreme-Go/application"
)

func main() {
	// data.Context.AutoMigrate(&entities.Account{})
	// data.Context.Create(&entities.Account{Code: "D42", Price: 110})
	application.Listen()
}
