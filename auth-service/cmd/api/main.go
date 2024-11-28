package main

import (
	"auth-service/cmd/api/db"
	"auth-service/cmd/api/models"
	"auth-service/cmd/api/repository"
	"auth-service/cmd/api/service"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)


func main() {
	dbConfigurations := db.DBConfig{
		Host: "pg-3c3beeec-nozamas-a08b.e.aivencloud.com",
		Port: 20674,
		User: "avnadmin",
		Password: "AVNS_IFRJnnnfQ7bB2BImnVu",
		Database: "defaultdb",
	}

	connection := db.GetConnection(dbConfigurations)

	userRepository := repository.UserRepository{
		Connection: connection,
	}

	userService := service.UserService{
		UserRepository: userRepository,
	}

	//Save User
	savingUserEmail := "sidath@gmail.com"
	userService.SaveUser(models.User{
		Id: fmt.Sprintf("%d", time.Now().UnixNano()),
		Email: savingUserEmail,
		Password: "sidath",
	})

	// Get Users
	isValidLogin := userService.IsValidUserLogin(savingUserEmail, "sidath")
	fmt.Println("isValidLogin", isValidLogin)

}