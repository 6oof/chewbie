package app

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

// Add your existing test functions here...

func TestRegisterRoutes(t *testing.T) {
	r := chi.NewRouter()
	registerRoutes(r)

	// Use httptest.NewServer instead of test.NewServer
	ts := httptest.NewServer(r)
	defer ts.Close()

	testCases := []struct {
		method   string
		path     string
		expected string
	}{
		{http.MethodGet, "/", "<!doctype html>"},
		{http.MethodPost, "/showcase-form", "<div class=\"mt-8\" id=\"result\" hx-swap=\"outerHTML\">"},
		// Add more cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.path, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, ts.URL+tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			// Check the response code
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, resp.StatusCode)
			}

			// Check if the expected content is present in the actual response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			if !strings.Contains(string(body), tc.expected) {
				t.Errorf("Expected content %s not found in the response body\n", tc.expected)
			}
		})
	}
}
