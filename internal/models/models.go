package models

type AnalyzeRequest struct {
	URL string `json:"url"`
}

type AnalyzeResponse struct {
	URL              string `json:"url"`
	StatusCode       int    `json:"statusCode"`
	ResponseTimeMs   int64  `json:"responseTimeMs"`
	Title            string `json:"title"`
	MetaDescription  string `json:"metaDescription"`
	H1Count          int    `json:"h1Count"`
	MissingAltImages int    `json:"missingAltImages"`
	WordCount        int    `json:"wordCount"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Code    string `json:"code"`
	Message string `json:"message"`
}
