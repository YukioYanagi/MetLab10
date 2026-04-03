package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingReturnsStaticJSON(t *testing.T) {
	router := newRouter()

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if got := rec.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
		t.Fatalf("expected content-type application/json; charset=utf-8, got %q", got)
	}

	if got := rec.Body.String(); got != `{"ok":true}` {
		t.Fatalf("expected body %q, got %q", `{"ok":true}`, got)
	}
}

func TestUnknownRouteReturns404(t *testing.T) {
	router := newRouter()

	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rec.Code)
	}
}

