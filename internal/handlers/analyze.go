package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"pagepulse/internal/models"
	"pagepulse/internal/parser"
	"pagepulse/internal/services"
)

func Analyze(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, &models.ErrorResponse{
			Error:   "METHOD_NOT_ALLOWED",
			Code:    "405",
			Message: "only POST is allowed",
		})
		return
	}

	var req models.AnalyzeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, &models.ErrorResponse{
			Error:   models.ErrInvalidURL,
			Code:    "400",
			Message: "invalid request body",
		})
		return
	}

	req.URL = strings.TrimSpace(req.URL)
	if req.URL == "" {
		writeError(w, http.StatusBadRequest, &models.ErrorResponse{
			Error:   models.ErrInvalidURL,
			Code:    "400",
			Message: "url field is required",
		})
		return
	}

	result, fetchErr := services.FetchPage(req.URL)
	if fetchErr != nil {
		code := 502
		switch fetchErr.Code {
		case "400":
			code = http.StatusBadRequest
		case "422":
			code = http.StatusUnprocessableEntity
		case "504":
			code = http.StatusGatewayTimeout
		}
		writeError(w, code, fetchErr)
		return
	}

	response := parser.Parse(result.StatusCode, result.ResponseTimeMs, req.URL, result.Body)
	writeJSON(w, http.StatusOK, response)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, e *models.ErrorResponse) {
	writeJSON(w, status, e)
}
