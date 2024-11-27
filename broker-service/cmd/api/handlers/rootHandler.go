package handlers

import (
	"broker/cmd/api/helpers"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	payload := helpers.JsonResponse{
		Error: false,
		Message: "Hit the broker",
	}
	helpers.CreateResponse(w, payload, http.StatusAccepted, make(map[string]string))
}
