# 🚀 PagePulse

A full-stack web application that analyzes any publicly accessible webpage and provides key performance and SEO insights. Built with **Go** for the backend and **React (Vite)** for the frontend.

> Built for the **Digital Heroes Software Development Training Task**.

---

## 🌐 Live Demo

- **Frontend:** https://page-pulse-nu.vercel.app/
- **Backend API:** https://pagepulse-production-72d3.up.railway.app
- **GitHub Repository:** https://github.com/VinayGadKerkar/PagePulse.git
- **Demo Video:** https://www.loom.com/share/472446d7a69741a591e509be880d5486

---

## ✨ Features

- Analyze any public webpage by URL
- Measure HTTP response time
- Display HTTP status code
- Extract page title
- Extract meta description
- Count `<h1>` tags
- Count words on the page
- Detect images missing `alt` attributes
- Robust error handling for invalid URLs, network failures, and non-HTML responses

---

## 🛠️ Tech Stack

### Backend

- Go
- net/http
- goquery
- Docker

### Frontend

- React
- Vite

### Deployment

- Railway (Backend)
- Vercel (Frontend)

---

## 🏗️ Project Structure

```text
pagepulse/
├── cmd/
│   └── server/
├── internal/
│   ├── handlers/
│   ├── middleware/
│   ├── parser/
│   ├── models/
│   └── services/
├── frontend/
├── Dockerfile
├── railway.toml
└── README.md
```

---

## 🧠 Architecture

```
React Frontend
       │
       ▼
Go REST API
       │
       ▼
HTTP Client
       │
       ▼
Target Website
       │
       ▼
HTML Parser (goquery)
       │
       ▼
JSON Response
```

---

## 📡 API

### POST `/api/analyze`

Analyzes a webpage.

### Request

```http
POST /api/analyze
Content-Type: application/json
```

```json
{
  "url": "https://example.com"
}
```

### Success Response

```json
{
  "url": "https://example.com",
  "statusCode": 200,
  "responseTimeMs": 127,
  "title": "Example Domain",
  "metaDescription": "...",
  "h1Count": 1,
  "wordCount": 145,
  "missingAltImages": 0
}
```

---

## ❌ Error Responses

Example:

```json
{
  "error": "INVALID_URL",
  "message": "Please provide a valid URL."
}
```

Other supported errors include:

- NETWORK_ERROR
- REQUEST_TIMEOUT
- NON_HTML_RESPONSE
- INTERNAL_SERVER_ERROR

---
---

## 💡 Design Decisions

### 1. Separation of Concerns

The project is divided into handlers, services, parser, middleware, and models to keep HTTP logic, business logic, and HTML parsing independent, making the codebase easier to maintain and test.

### 2. Server-side HTML Parsing

The backend performs all webpage analysis using `goquery` instead of relying on the frontend. This keeps the API reusable and avoids browser-specific limitations.

### 3. Consistent Error Responses

All failures return structured JSON responses with meaningful error codes such as `INVALID_URL`, `NETWORK_ERROR`, and `NON_HTML_RESPONSE`, allowing the frontend to handle errors consistently.

---

## 🚀 Running Locally

### Clone

```bash
git clone https://github.com/VinayGadKerkar/PagePulse.git
cd pagepulse
```

### Backend

```bash
go mod download
go run ./cmd/server
```

Runs on:

```
http://localhost:8080
```

### Frontend

```bash
cd frontend
npm install
npm run dev
```

Runs on:

```
http://localhost:5173
```

---

## 🐳 Docker

Build:

```bash
docker build -t pagepulse .
```

Run:

```bash
docker run -p 8080:8080 pagepulse
```

---


## 🧪 Testing

Run backend tests:

```bash
go test ./...
```

---

## 🔮 Future Improvements

- Lighthouse integration
- Open Graph metadata extraction
- Sitemap detection
- Robots.txt analysis
- Canonical URL detection
- Performance history
- Authentication & saved reports
- Batch URL analysis

---

## 🤖 AI Usage

AI tools (**ChatGPT** and **Claude**) were used throughout the project to:

- Brainstorm the overall project architecture
- Generate and iterate on portions of the backend and frontend code
- Review and refactor Go code for readability and maintainability
- Discuss error-handling approaches and API design
- Improve the README documentation
- Refine UI copy and project structure

After using AI-generated suggestions, I reviewed the code, modified it to fit the project requirements, debugged implementation issues, fixed deployment and CORS problems, wrote and verified the tests, and made the final implementation and deployment decisions myself. All submitted code was manually tested and validated before submission.

---

## 📄 License

This project was built as part of the **Digital Heroes Software Development Training Task**.

---

## 🙏 Acknowledgements

Special thanks to **Digital Heroes** for designing this practical backend engineering assignment.

https://digitalheroesco.com
