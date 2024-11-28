package service

import (
	"auth-service/cmd/api/models"
	"auth-service/cmd/api/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func (userService *UserService) SaveUser(user models.User) models.User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	user.Password = string(hashedPassword)
	return userService.UserRepository.SaveUser(user)
}

func (userService *UserService) IsValidUserLogin(email string, password string) bool {
	foundUserByEmail := userService.UserRepository.FindUserByEmail(email)
	err := bcrypt.CompareHashAndPassword([]byte(foundUserByEmail.Password), []byte(password))
	return err == nil

}