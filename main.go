package main

import (
	"gincrud/models"
	"gincrud/routes"
)

func main(){

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)

	r.Run()

}