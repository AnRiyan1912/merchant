package main

import (
	"rename.com/andreriyant/go-crud/initializers"
	"rename.com/andreriyant/go-crud/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvirontmentVariable()
}


func main() {
	initializers.DB.AutoMigrate(&models.User{})
}