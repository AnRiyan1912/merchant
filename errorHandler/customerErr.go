package errorhandler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"rename.com/andreriyant/go-crud/services"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func HandleError(c *gin.Context, err error) {
	var status  int
	var message string

	switch err {
	case services.ErrUsernameTaken:
		status = http.StatusConflict
		message = "Username already taken"
	case services.ErrUserNotFound:
		status = http.StatusNotFound
		message = "User not found"
	case services.ErrInvalidCredentials:
		status = http.StatusUnauthorized
		message = "Invalid credentials"
	default:
		status = http.StatusInternalServerError
		message = "Internal Server Error"
	}

	c.JSON(status, ErrorResponse{Message: message, Status: status})
}

// Jangan lupa untuk mendefinisikan variabel error khusus di sini
var (
	ErrUsernameTaken    = errors.New("username is already taken")
	ErrUserNotFound     = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
