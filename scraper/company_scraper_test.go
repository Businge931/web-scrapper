package scraper

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestReadCompanyNames(t *testing.T) {
	// Prepare a temporary directory and file for testing
	tempDir := "companies-list"
	tempFile := tempDir + "/input.txt"
	expectedNames := []string{"stanbic bank", "Pearl Technologies", "clinicpesa"}

	// Ensure the directory exists
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}

	// Create the input file with expected company names
	file, err := os.Create(tempFile)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up the temp directory after the test
	defer file.Close()

	// Write the expected company names to the file
	for _, name := range expectedNames {
		if _, err := file.WriteString(name + "\n"); err != nil {
			t.Fatalf("Failed to write to temp file: %v", err)
		}
	}

	// Call the function under test
	actualNames, err := ReadCompanyNames(tempFile)
	if err != nil {
		t.Fatalf("ReadCompanyNames returned an error: %v", err)
	}

	// Compare the actual output with the expected output
	if !reflect.DeepEqual(actualNames, expectedNames) {
		t.Errorf("Expected %v, but got %v", expectedNames, actualNames)
	}
}

func TestGetSearchResults(t *testing.T) {
	// Mock response data
	mockResponse := SerpAPIResponse{
		Organic: []struct {
			Link string `json:"link"`
		}{
			{Link: "https://google.serper.dev/search"},
		},
	}

	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Encode the mock response as JSON and write it to the response writer
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	// Create an http.Client that directs requests to the mock server
	client := server.Client()

	// Call the function under test with the mock client
	companyName := "stanbic bank"
	result, err := GetSearchResults(client, companyName)
	if err != nil {
		t.Fatalf("GetSearchResults returned an error: %v", err)
	}

	// Expected URL from the mock response
	expectedURL := "https://www.stanbicbank.co.ug/uganda/personal"

	// Check if the result matches the expected URL
	if result != expectedURL {
		t.Errorf("Expected %s, but got %s", expectedURL, result)
	}
}
