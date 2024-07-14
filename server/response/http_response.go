package response

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Status  int         `json:"status,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func RespondWithJson(w http.ResponseWriter, message string, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := HttpResponse{
		Message: message,
		Data:    data,
		Status:  statusCode,
	}
	json.NewEncoder(w).Encode(response)

}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := HttpResponse{
		Error:  message,
		Status: statusCode,
	}
	json.NewEncoder(w).Encode(response)
}
