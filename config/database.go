package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	var DB_USERNAME = "root"
	var DB_PASSWORD = "root"
	var DB_HOST = "localhost"
	var DB_PORT = "8889"
	var DB_NAME = "gdas"

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("CONECTED TO DATABASE")
	return db, nil
}
