package main

import (
	"github.com/gin-gonic/gin"
	"rename.com/andreriyant/go-crud/controllers"
	"rename.com/andreriyant/go-crud/initializers"
)

func init() {
	initializers.LoadEnvirontmentVariable()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)


	r.Run()
}
