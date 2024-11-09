package validation

import (
	"encoding/json"
	"net/http"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
)

func RequiredValidate(w http.ResponseWriter, massage string)  {
	validate := model.Response{
		Status: false,
		Message: massage,
	}

	json.NewEncoder(w).Encode(validate)
}