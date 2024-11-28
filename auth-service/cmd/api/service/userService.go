package service

import (
	"auth-service/cmd/api/models"
	"auth-service/cmd/api/repository"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: *repository.NewUserRepository(),
	}
}

func (userService *UserService) SaveUser(user models.User) models.User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}

	user.Password = string(hashedPassword)
	return userService.userRepository.SaveUser(user)
}

func (userService *UserService) IsValidUserLogin(email string, password string) bool {
	foundUserByEmail := userService.userRepository.FindUserByEmail(email)
	log.Println("Found user with email:", foundUserByEmail)
	err := bcrypt.CompareHashAndPassword([]byte(foundUserByEmail.Password), []byte(password))
	return err == nil

}