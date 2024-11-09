package validation

import (
	"encoding/json"
	"net/http"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
)

func BadResponse(w http.ResponseWriter, message string , statusCode int)  {
	response := model.Response {
		Status: false,
		StatusCode: statusCode,
		Message: message,
		Data: nil,
	}

	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func OKResponse(w http.ResponseWriter, message string , data any)  {
	response := model.Response {
		Status: true,
		StatusCode: http.StatusOK,
		Message: message,
		Data: data,
	}

	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

