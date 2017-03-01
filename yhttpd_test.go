package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeRequestGet(t *testing.T) {
	for _, url := range []string {"/", "/foo/bar"} {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Errorf("Expected to not throw error: %s", err)
		}
		w := httptest.NewRecorder()
		serveRequest(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("Expected response code %d, got: %d", http.StatusOK, w.Code)
		}
		if w.Body.String() != "OK" {
			t.Errorf("Expected response  'OK', got: %s", w.Body.String())
		}
		h := w.HeaderMap.Get("Content-Type")
		if h != "text/plain; charset=utf-8" {
			t.Errorf("Expected header 'Content-Type' to be 'text/plain; charset=utf8', got: %s", h)
		}
	}
}

func TestServeRequestPost(t *testing.T) {
	for _, url := range []string {"/", "/foo/bar"} {
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			t.Errorf("Expected to not throw error: %s", err)
		}
		w := httptest.NewRecorder()
		serveRequest(w, req)
		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected response code %d, got: %d", http.StatusMethodNotAllowed, w.Code)
		}
	}
}

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	go listen(8099)
	rsp, err := http.Get("http://localhost:8099/blubb")
	if err != nil {
		t.Errorf("Expected to not throw error: %s", err)
	}
	if rsp.StatusCode != http.StatusOK {
		t.Errorf("Expected response code %d, got: %d", http.StatusOK, rsp.StatusCode)
	}
}
