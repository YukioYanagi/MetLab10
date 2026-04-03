package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPingEndpoint(t *testing.T) {
	router := newRouter()

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if got := rec.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
		t.Fatalf("unexpected content-type: %q", got)
	}
}

func TestOpenAPIEndpoint(t *testing.T) {
	router := newRouter()

	req := httptest.NewRequest(http.MethodGet, "/openapi.json", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if body == "" {
		t.Fatal("openapi response must not be empty")
	}
	if !strings.Contains(body, `"openapi":"3.0.3"`) {
		t.Fatalf("openapi version not found in body: %s", body)
	}
	if !strings.Contains(body, `"/ping"`) {
		t.Fatalf("ping path not found in body: %s", body)
	}
}

