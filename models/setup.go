package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"

)

func SetupDB() *gorm.DB {
	USER := "gocrud"
	PASS := "MySql2019!"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "gocrud"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	
	

	if err != nil {
		panic(err.Error())
	}
	
	return db
}