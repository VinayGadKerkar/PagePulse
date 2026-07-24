package services

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"pagepulse/internal/models"
)

type FetchResult struct {
	StatusCode     int
	ResponseTimeMs int64
	Body           string
}

func FetchPage(rawURL string) (*FetchResult, *models.ErrorResponse) {
	// Validate URL
	parsed, err := url.ParseRequestURI(rawURL)
	if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		return nil, &models.ErrorResponse{
			Error:   models.ErrInvalidURL,
			Code:    "400",
			Message: "the provided URL is not valid",
		}
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	start := time.Now()
	resp, err := client.Get(rawURL)
	elapsed := time.Since(start).Milliseconds()

	if err != nil {
		return nil, classifyError(err)
	}
	defer resp.Body.Close()

	// Guard against ISP DNS hijacking — they return 200 with their own page
	if resp.StatusCode == 200 && resp.Request.URL.Host != parsed.Host {
		return nil, &models.ErrorResponse{
			Error:   models.ErrDNSFailure,
			Code:    "502",
			Message: "could not resolve host",
		}
	}

	// Check content type
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return nil, &models.ErrorResponse{
			Error:   models.ErrNonHTML,
			Code:    "422",
			Message: fmt.Sprintf("expected text/html but got %s", contentType),
		}
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &models.ErrorResponse{
			Error:   models.ErrFetchError,
			Code:    "502",
			Message: "failed to read response body",
		}
	}

	return &FetchResult{
		StatusCode:     resp.StatusCode,
		ResponseTimeMs: elapsed,
		Body:           string(bodyBytes),
	}, nil
}

func classifyError(err error) *models.ErrorResponse {
	// Timeout
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return &models.ErrorResponse{
			Error:   models.ErrTimeout,
			Code:    "504",
			Message: "the request timed out",
		}
	}

	errStr := err.Error()
	// DNS failure
	if strings.Contains(errStr, "no such host") ||
		strings.Contains(errStr, "dns") ||
		strings.Contains(errStr, "lookup") ||
		strings.Contains(errStr, "dial") {
		return &models.ErrorResponse{
			Error:   models.ErrDNSFailure,
			Code:    "502",
			Message: "could not resolve host",
		}
	}

	return &models.ErrorResponse{
		Error:   models.ErrFetchError,
		Code:    "502",
		Message: errStr,
	}
}
