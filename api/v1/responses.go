package v1

import (
	"encoding/json"
	"net/http"
)

//JsonResponse Type
type JsonResponse struct {
	// Reserved field to add some meta information to the API response
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

//JsonErrorResponse Type
type JsonErrorResponse struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Error   string `json:"message"`
}

// Writes the response as a standard JSON response with StatusOK
func WriteOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&JsonResponse{Success: true, Status: http.StatusOK, Data: m}); err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "internal server error")
	}
}

// Writes the error response as a Standard API JSON response with a response code
func WriteErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.
		NewEncoder(w).
		Encode(&JsonErrorResponse{Success: false, Status: errorCode, Error: errorMsg})
}
