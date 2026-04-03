package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testLogger struct {
	lines []string
}

func (t *testLogger) Printf(format string, v ...any) {
	t.lines = append(t.lines, fmt.Sprintf(format, v...))
}

func TestPingEndpointReturns200(t *testing.T) {
	log := &testLogger{}
	router := newRouter(log)

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if got := rec.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
		t.Fatalf("unexpected content type: %q", got)
	}
	if rec.Body.String() != `{"ok":true}` {
		t.Fatalf("unexpected body: %s", rec.Body.String())
	}
}

func TestLoggingMiddlewareWritesRequestData(t *testing.T) {
	log := &testLogger{}
	router := newRouter(log)

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if len(log.lines) == 0 {
		t.Fatal("expected at least one log line")
	}
	line := log.lines[len(log.lines)-1]
	if !strings.Contains(line, "method=GET") {
		t.Fatalf("missing method in log line: %s", line)
	}
	if !strings.Contains(line, "path=/ping") {
		t.Fatalf("missing path in log line: %s", line)
	}
	if !strings.Contains(line, "status=200") {
		t.Fatalf("missing status in log line: %s", line)
	}
	if !strings.Contains(line, "duration_ms=") {
		t.Fatalf("missing duration in log line: %s", line)
	}
}

