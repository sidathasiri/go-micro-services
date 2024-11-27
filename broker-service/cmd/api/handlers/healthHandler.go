package handlers

import (
	"broker/cmd/api/helpers"
	"net/http"
)

func GetHealthHandler(w http.ResponseWriter, r *http.Request) {
	payload := helpers.JsonResponse{
		Error: false,
		Message: "Working",
	}
	helpers.CreateResponse(w, payload, http.StatusAccepted, map[string]string{})
}
