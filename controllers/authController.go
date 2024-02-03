package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errorhandler "rename.com/andreriyant/go-crud/errorHandler"
	"rename.com/andreriyant/go-crud/models"
	"rename.com/andreriyant/go-crud/services"
)

func Register(c *gin.Context) {
  var user models.User
  var person models.Person

  if err := services.RegisterUser(&user, &person); err != nil {
	errorhandler.HandleError(c, err)
	return
  }
  c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})

}

func Login(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	token, err := services.AuthenticateUser(loginRequest.Username, loginRequest.Password)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
