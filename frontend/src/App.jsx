import { useState } from "react";

const STATUS_COLORS = {
  success: "#22c55e",
  error: "#ef4444",
  warning: "#f59e0b",
};

function getStatusColor(code) {
  if (code >= 200 && code < 300) return STATUS_COLORS.success;
  if (code >= 400) return STATUS_COLORS.error;
  return STATUS_COLORS.warning;
}

function MetricCard({ label, value, highlight }) {
  return (
    <div style={{
      background: "#111827",
      border: "1px solid #1f2937",
      borderRadius: "8px",
      padding: "16px 20px",
      display: "flex",
      flexDirection: "column",
      gap: "4px",
    }}>
      <span style={{ color: "#6b7280", fontSize: "12px", textTransform: "uppercase", letterSpacing: "0.05em" }}>
        {label}
      </span>
      <span style={{ color: highlight || "#f9fafb", fontSize: "22px", fontWeight: "600" }}>
        {value}
      </span>
    </div>
  );
}

export default function App() {
  const API = import.meta.env.VITE_API_URL;
  const [url, setUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState(null);
  const [error, setError] = useState(null);

  async function analyze() {
    const trimmed = url.trim();
    if (!trimmed) return;

    setLoading(true);
    setResult(null);
    setError(null);

    try {
      const res = await fetch(`${API}/analyze`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ url: trimmed }),
      });

      const data = await res.json();

      if (!res.ok) {
        setError(data);
      } else {
        setResult(data);
      }
    } catch (e) {
      setError({ error: "NETWORK_ERROR", message: "Could not reach the server." });
    } finally {
      setLoading(false);
    }
  }

  function handleKeyDown(e) {
    if (e.key === "Enter") analyze();
  }

  return (
    <div style={{
      minHeight: "100vh",
      background: "#030712",
      color: "#f9fafb",
      fontFamily: "'Inter', system-ui, sans-serif",
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
      padding: "60px 20px",
    }}>

      {/* Header */}
      <div style={{ textAlign: "center", marginBottom: "48px" }}>
        <h1 style={{ fontSize: "36px", fontWeight: "700", margin: 0, letterSpacing: "-0.02em" }}>
          Page<span style={{ color: "#6366f1" }}>Pulse</span>
        </h1>
        <p style={{ color: "#6b7280", marginTop: "8px", fontSize: "15px" }}>
          Analyse any webpage — status, performance, and SEO signals at a glance.
        </p>
      </div>

      {/* Input */}
      <div style={{ display: "flex", gap: "8px", width: "100%", maxWidth: "600px" }}>
        <input
          type="text"
          placeholder="https://example.com"
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          onKeyDown={handleKeyDown}
          style={{
            flex: 1,
            background: "#111827",
            border: "1px solid #1f2937",
            borderRadius: "8px",
            padding: "12px 16px",
            color: "#f9fafb",
            fontSize: "15px",
            outline: "none",
          }}
        />
        <button
          onClick={analyze}
          disabled={loading}
          style={{
            background: loading ? "#4338ca" : "#6366f1",
            color: "#fff",
            border: "none",
            borderRadius: "8px",
            padding: "12px 24px",
            fontSize: "15px",
            fontWeight: "600",
            cursor: loading ? "not-allowed" : "pointer",
            transition: "background 0.2s",
          }}
        >
          {loading ? "Analysing..." : "Analyse"}
        </button>
      </div>

      {/* Error */}
      {error && (
        <div style={{
          marginTop: "32px",
          width: "100%",
          maxWidth: "600px",
          background: "#1c0a0a",
          border: "1px solid #7f1d1d",
          borderRadius: "8px",
          padding: "16px 20px",
        }}>
          <p style={{ color: "#ef4444", fontWeight: "600", margin: 0 }}>{error.error}</p>
          <p style={{ color: "#9ca3af", margin: "4px 0 0" }}>{error.message}</p>
        </div>
      )}

      {/* Results */}
      {result && (
        <div style={{ marginTop: "32px", width: "100%", maxWidth: "600px" }}>

          {/* URL + title */}
          <div style={{
            background: "#111827",
            border: "1px solid #1f2937",
            borderRadius: "8px",
            padding: "16px 20px",
            marginBottom: "12px",
          }}>
            <p style={{ margin: 0, color: "#6b7280", fontSize: "12px", textTransform: "uppercase", letterSpacing: "0.05em" }}>Page</p>
            <p style={{ margin: "4px 0 0", fontWeight: "600", fontSize: "18px" }}>{result.title || "—"}</p>
            <p style={{ margin: "4px 0 0", color: "#6b7280", fontSize: "13px", wordBreak: "break-all" }}>{result.url}</p>
            {result.metaDescription && (
              <p style={{ margin: "8px 0 0", color: "#9ca3af", fontSize: "14px" }}>{result.metaDescription}</p>
            )}
          </div>

          {/* Metrics grid */}
          <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: "12px" }}>
            <MetricCard
              label="Status"
              value={result.statusCode}
              highlight={getStatusColor(result.statusCode)}
            />
            <MetricCard
              label="Response Time"
              value={`${result.responseTimeMs}ms`}
              highlight={result.responseTimeMs < 500 ? STATUS_COLORS.success : STATUS_COLORS.warning}
            />
            <MetricCard label="H1 Tags" value={result.h1Count} />
            <MetricCard label="Word Count" value={result.wordCount} />
            <MetricCard
              label="Images Missing Alt"
              value={result.missingAltImages}
              highlight={result.missingAltImages > 0 ? STATUS_COLORS.warning : STATUS_COLORS.success}
            />
          </div>
        </div>
      )}
      <footer
        style={{
          marginTop: "60px",
          color: "#6b7280",
          fontSize: "13px",
        }}
      >
        Built for{" "}
        <a
          href="https://digitalheroesco.com"
          target="_blank"
          rel="noopener noreferrer"
          style={{
            color: "#6366f1",
            textDecoration: "none",
          }}
        >
          Digital Heroes Training Task
        </a>
      </footer>
    </div>
  );
}