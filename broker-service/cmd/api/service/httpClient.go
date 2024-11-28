package service

import (
	"broker/cmd/api/helpers"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type HttpClient struct {}

func (httpClient *HttpClient) Post (requestPath string, body io.Reader) (helpers.JsonResponse, error, int) {
	request, _ := http.NewRequest("POST", requestPath, body)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Failed to send request:", err)
		return helpers.JsonResponse{}, err, http.StatusInternalServerError
	}
	
	defer response.Body.Close()

	// Read the response from the server
	var serverResponse helpers.JsonResponse
	err = json.NewDecoder(response.Body).Decode(&serverResponse)
	if err != nil {
		log.Println("Failed to decode server response:", err)
		return helpers.JsonResponse{}, err, http.StatusInternalServerError
	}

	return serverResponse, nil, http.StatusOK
}

