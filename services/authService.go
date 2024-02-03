package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"rename.com/andreriyant/go-crud/initializers"
	"rename.com/andreriyant/go-crud/models"
	"rename.com/andreriyant/go-crud/utils"
)

var (
	ErrUsernameTaken      = errors.New("username is already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func RegisterUser(request *models.RegisterRequest) error {
	user := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	person := models.Person{
		Fullname: request.Person.Fullname,
		Email:    request.Person.Email,
		Address:  request.Person.Address,
	}

	if err := initializers.DB.Create(&person).Error; err != nil {
		return err
	}

	if person.Id == 0 {
		return errors.New("failed to get Person ID")
	}

	user.PersonID = person.Id

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if err := initializers.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func AuthenticateUser(username, password string) (string, error) {

    user := models.User{}
    if err := initializers.DB.Where("username = ?", username).First(&user).Error; err != nil {
     
        return "", ErrUserNotFound
    }
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {

        return "", ErrInvalidCredentials
    }
    token, err := utils.GenerateToken(username)
    if err != nil {
        return "", err
    }

    return token, nil
}