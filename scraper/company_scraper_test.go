package scraper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestReadCompanyNames(t *testing.T) {
	names, err := ReadCompanyNames("data/input.txt")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(names) != 10 {
		t.Fatalf("expected 3 names, got %d", len(names))
	}

	expected := "Kampala Digital Hub"
	if names[0] != expected {
		t.Errorf("expected %s, got %s", expected, names[0])
	}
}

var baseURL = "https://google.serper.dev/search"

func TestGetSearchResults(t *testing.T) {
	// Define a mock response for the API
	mockResponse := `{
		"organic": [
			{
				"link": "https://www.stanbicbank.co.ug/uganda/personal"
			}
		]
	}`

	// Start a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("q") == "stanbic bank" {
			fmt.Fprintln(w, mockResponse)
		} else {
			http.Error(w, "not found", http.StatusNotFound)
		}
	}))
	defer mockServer.Close()

	// Temporarily replace the base URL with the mock server's URL
	originalBaseURL := "https://google.serper.dev/search"
	defer func() { baseURL = originalBaseURL }()
	baseURL = mockServer.URL

	// Test case: successful search result
	t.Run("Successful search result", func(t *testing.T) {
		url, err := GetSearchResults("stanbic bank")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		expectedURL := "https://www.stanbicbank.co.ug/uganda/personal"
		if url != expectedURL {
			t.Errorf("expected %v, got %v", expectedURL, url)
		}
	})

	// Test case: no results found
	t.Run("No results found", func(t *testing.T) {
		_, err := GetSearchResults("unknown company")
		if err == nil {
			t.Fatalf("expected an error, got none")
		}
		expectedError := "no results found for unknown company"
		if err.Error() != expectedError {
			t.Errorf("expected error %v, got %v", expectedError, err)
		}
	})

	// Test case: API key not set
	t.Run("API key not set", func(t *testing.T) {
		os.Setenv("SERPAPI_KEY", "")
		_, err := GetSearchResults("stanbic bank")
		if err == nil {
			t.Fatalf("expected an error, got none")
		}
		expectedError := "SERPAPI_KEY not set"
		if err.Error() != expectedError {
			t.Errorf("expected error %v, got %v", expectedError, err)
		}
	})
}
