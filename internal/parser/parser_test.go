package parser_test

import (
	"testing"

	"pagepulse/internal/parser"
)

const happyHTML = `
<html>
  <head>
    <title>Hello World</title>
    <meta name="description" content="A test page" />
  </head>
  <body>
    <h1>Main Heading</h1>
    <img src="a.jpg" alt="good image" />
    <img src="b.jpg" />
    <img src="c.jpg" alt="" />
    <p>One two three four five</p>
  </body>
</html>`

func TestParse_HappyPath(t *testing.T) {
	result := parser.Parse(200, 100, "https://example.com", happyHTML)

	if result.Title != "Hello World" {
		t.Errorf("expected title 'Hello World', got '%s'", result.Title)
	}
	if result.MetaDescription != "A test page" {
		t.Errorf("expected meta 'A test page', got '%s'", result.MetaDescription)
	}
	if result.H1Count != 1 {
		t.Errorf("expected 1 h1, got %d", result.H1Count)
	}
	if result.MissingAltImages != 2 {
		t.Errorf("expected 2 missing alts, got %d", result.MissingAltImages)
	}
	if result.WordCount != 7 {
		t.Errorf("expected 7 words, got %d", result.WordCount)
	}
}

func TestParse_MissingTitle(t *testing.T) {
	html := `<html><body><h1>No title here</h1></body></html>`
	result := parser.Parse(200, 50, "https://example.com", html)
	if result.Title != "" {
		t.Errorf("expected empty title, got '%s'", result.Title)
	}
}

func TestParse_EmptyHTML(t *testing.T) {
	result := parser.Parse(200, 50, "https://example.com", "")
	if result.H1Count != 0 || result.WordCount != 0 {
		t.Errorf("expected zeros on empty HTML")
	}
}

func TestParse_MultipleH1(t *testing.T) {
	html := `<html><body><h1>One</h1><h1>Two</h1><h1>Three</h1></body></html>`
	result := parser.Parse(200, 50, "https://example.com", html)
	if result.H1Count != 3 {
		t.Errorf("expected 3 h1s, got %d", result.H1Count)
	}
}
