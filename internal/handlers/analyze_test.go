package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"pagepulse/internal/handlers"
	"pagepulse/internal/models"
)

func TestAnalyzeHandler_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/analyze", nil)
	w := httptest.NewRecorder()

	handlers.Analyze(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", w.Code)
	}
}

func TestAnalyzeHandler_EmptyBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/analyze", bytes.NewBufferString(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handlers.Analyze(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}

	var resp models.ErrorResponse
	json.NewDecoder(w.Body).Decode(&resp)
	if resp.Error != models.ErrInvalidURL {
		t.Errorf("expected INVALID_URL, got %s", resp.Error)
	}
}

func TestAnalyzeHandler_InvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/analyze", bytes.NewBufferString(`not json`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handlers.Analyze(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestAnalyzeHandler_InvalidURL(t *testing.T) {
	body := `{"url":"not-a-url"}`
	req := httptest.NewRequest(http.MethodPost, "/api/analyze", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handlers.Analyze(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}
