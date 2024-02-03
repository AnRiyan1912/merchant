package controllers

import (
	// "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"rename.com/andreriyant/go-crud/initializers"

	// "gorm.io/gorm"
	"rename.com/andreriyant/go-crud/models"
)

func Index(c *gin.Context){
   var users []models.User
   initializers.DB.Find(&users)
   c.JSON(http.StatusOK, gin.H{"products": users})

}
func Show(c *gin.Context){
	
}
func Update(c *gin.Context){
	
}
func Create(c *gin.Context){
	
}
func Delete(c *gin.Context){
	
}