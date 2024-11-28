package helpers

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}

func CreateResponse (w http.ResponseWriter, payload JsonResponse, statusCode int, headers map[string]string) error {
	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-type", "application/json")
	if(len(headers) > 0) {
		for key, value := range(headers) {
			w.Header().Set(key, value)
		}
	}
	w.WriteHeader(statusCode)
	_, err := w.Write(out)
	if err != nil {
		return err
	}
	return nil
}