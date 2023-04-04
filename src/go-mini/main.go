package main

import (
	"Emc002/go-mini/models"
	"Emc002/go-mini/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run()
}
