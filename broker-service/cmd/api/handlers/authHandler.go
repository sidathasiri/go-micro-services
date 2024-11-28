package handlers

import (
	"broker/cmd/api/helpers"
	"broker/cmd/api/service"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type AuthRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	httpClient := service.HttpClient{}
	receivedCredentials := AuthRequest{}
	err := json.NewDecoder(r.Body).Decode(&receivedCredentials)
	if err != nil {
		log.Println("Invalid request body")
		responseBody := helpers.JsonResponse{
			Error: true,
			Message: "Invalid request",
		}
		helpers.CreateResponse(w, responseBody, http.StatusBadRequest, make(map[string]string))
		return
	}
	payloadBytes, _ := json.Marshal(receivedCredentials)
	response, err, statusCode := httpClient.Post("http://localhost:4000/auth/login", bytes.NewReader(payloadBytes))
	if err != nil {
		responseBody := helpers.JsonResponse{
			Error: true,
			Message: err.Error(),
		}
		helpers.CreateResponse(w, responseBody, statusCode, make(map[string]string))
		return
	}

	responseBody := helpers.JsonResponse{
		Error: response.Error,
		Message: response.Message,
	}
	helpers.CreateResponse(w, responseBody, statusCode, make(map[string]string))
}
