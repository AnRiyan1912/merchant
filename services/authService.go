package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"rename.com/andreriyant/go-crud/initializers"
	"rename.com/andreriyant/go-crud/models"
	"rename.com/andreriyant/go-crud/utils"
)

var (
    ErrUsernameTaken   = errors.New("username is already taken")
    ErrUserNotFound    = errors.New("user not found")
    ErrInvalidCredentials = errors.New("invalid credentials")
)



func RegisterUser(user *models.User, person *models.Person) error {
	existingUser := models.User{}
	if err := initializers.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return  ErrUsernameTaken
	}
	if err := initializers.DB.Create(person).Error; err != nil {
		return err
	}
	user.PersonID = person.Id
	if err := initializers.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func AuthenticateUser(username, password string) (string, error) {
    // Temukan pengguna berdasarkan nama pengguna
    user := models.User{}
    if err := initializers.DB.Where("username = ?", username).First(&user).Error; err != nil {
        // Pengguna tidak ditemukan
        return "", ErrUserNotFound
    }

    // Verifikasi kata sandi
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        // Kata sandi tidak cocok
        return "", ErrInvalidCredentials
    }

    // Kata sandi cocok, hasilkan token JWT
    token, err := utils.GenerateToken(username)
    if err != nil {
        return "", err
    }

    return token, nil
}