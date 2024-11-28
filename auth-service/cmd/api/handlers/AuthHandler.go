package handlers

import (
	"auth-service/cmd/api/db"
	"auth-service/cmd/api/dto"
	"auth-service/cmd/api/helpers"
	"auth-service/cmd/api/repository"
	"auth-service/cmd/api/service"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)



func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var dbPort, _ = strconv.ParseUint(os.Getenv("DB_PORT"), 10, 16)

	var dbConfigurations = db.DBConfig{
		Host: os.Getenv("DB_HOST"),
		Port: uint(dbPort),
		User: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	var connection = db.GetConnection(dbConfigurations)

	var userRepository = repository.UserRepository{
		Connection: connection,
	}

	var userService = service.UserService{
		UserRepository: userRepository,
	}
	receivedCredentials := dto.Credentials{}
	err := json.NewDecoder(r.Body).Decode(&receivedCredentials)
	log.Println("Request received to authenticate with email:", receivedCredentials.Email)
	if err != nil {
		log.Println("Invalid request body")
		responseBody := helpers.JsonResponse{
			Error: true,
			Message: "Invalid request",
		}
		helpers.CreateResponse(w, responseBody, http.StatusBadRequest, make(map[string]string))
		return
	}
	var isValidLogin bool = userService.IsValidUserLogin(receivedCredentials.Email, receivedCredentials.Password) 
	if isValidLogin {
		log.Println(("Valid credentials found"))
		responseBody := helpers.JsonResponse{
			Error: false,
			Message: "Successfully authenticated",
		}
		helpers.CreateResponse(w, responseBody, http.StatusOK, make(map[string]string))
		return
	} else {
		log.Println("Invalid credentials found")
		responseBody := helpers.JsonResponse{
			Error: true,
			Message: "Invalid credentials",
		}
		helpers.CreateResponse(w, responseBody, http.StatusForbidden, make(map[string]string))
		return
	}
}
