package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/KannelRetry/response"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	data, _ := json.Marshal(response.ErrorResponse{Success: false, Error: response.ErrorResponseData{Code: code, Message: message}})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
