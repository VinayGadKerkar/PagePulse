package parser

import (
	"strings"

	"pagepulse/internal/models"

	"github.com/PuerkitoBio/goquery"
)

// Parse accepts an HTTP response body as a string and returns
// structured page metadata. It has zero knowledge of HTTP —
// that separation is what makes it fully unit-testable.
func Parse(statusCode int, responseTimeMs int64, url string, body string) models.AnalyzeResponse {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return models.AnalyzeResponse{
			URL:            url,
			StatusCode:     statusCode,
			ResponseTimeMs: responseTimeMs,
		}
	}

	title := strings.TrimSpace(doc.Find("title").First().Text())

	metaDesc := ""
	doc.Find("meta").Each(func(_ int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		if strings.ToLower(name) == "description" {
			metaDesc, _ = s.Attr("content")
		}
	})

	h1Count := doc.Find("h1").Length()

	missingAlt := 0
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		alt, exists := s.Attr("alt")
		if !exists || strings.TrimSpace(alt) == "" {
			missingAlt++
		}
	})

	wordCount := countWords(doc.Find("body").Text())

	return models.AnalyzeResponse{
		URL:              url,
		StatusCode:       statusCode,
		ResponseTimeMs:   responseTimeMs,
		Title:            title,
		MetaDescription:  metaDesc,
		H1Count:          h1Count,
		MissingAltImages: missingAlt,
		WordCount:        wordCount,
	}
}

func countWords(text string) int {
	fields := strings.Fields(text)
	return len(fields)
}
