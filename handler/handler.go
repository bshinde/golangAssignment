package handler

import (
	"encoding/json"
	"golangAssignment/model"
	"golangAssignment/service"
	"net/http"
)

// Handler handles incoming HTTP requests
type Handler struct {
	Service *service.Service
}

// NewHandler creates a new instance of the handler
func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

// HandleRequest is the handler function for the POST /validate endpoint
func (h *Handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}
	var payload model.RequestPayload

	// Parse the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		// If the request body is invalid, return a 400 Bad Request with an error message
		http.Error(w, "Invalid JSON format or request body", http.StatusBadRequest)
		return
	}

	// Validate the payload using the service
	if err := h.Service.ValidatePayload(payload); err != nil {
		// If validation fails, return the validation error message with a 400 Bad Request status
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return success message if validation passes
	response := map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Request validated successfully",
	}

	// Set the response header for content type to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and send the success response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Handle any encoding errors and return an internal server error if needed
		http.Error(w, "Failed to generate response", http.StatusInternalServerError)
	}
}
