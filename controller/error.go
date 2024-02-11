package controller

import (
	"encoding/json"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func writeErrorResponse(msg string, w http.ResponseWriter) {
	logger.Error(msg)

	w.WriteHeader(http.StatusInternalServerError)

	resp := ErrorResponse{Message: msg}
	json.NewEncoder(w).Encode(resp)
}
