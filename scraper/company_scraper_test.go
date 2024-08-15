package scraper

import (
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




