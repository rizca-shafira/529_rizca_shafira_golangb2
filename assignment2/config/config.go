package config

import (
	"ass2/structs"
	"fmt"

	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/orders_by?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Connected")

	// auto migrate models to database table
	db.AutoMigrate(structs.Item{}, structs.Order{})
	return db
}
