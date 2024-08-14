package scraper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-querystring/query"
)

// SerpAPI response struct
type SerpAPIResponse struct {
	Organic []struct {
		Link string `json:"link"`
	} `json:"organic"`
}

func ReadCompanyNames(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var companyNames []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		companyNames = append(companyNames, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return companyNames, nil
}

func GetSearchResults(companyName string) (string, error) {
	os.Setenv("SERPAPI_KEY", "0eb5aec35da6593d1993b1573558d3b5f8b0a37c")
	apiKey := os.Getenv("SERPAPI_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("SERPAPI_KEY not set")
	}

	baseURL := "https://google.serper.dev/search"
	params := struct {
		Query  string `url:"q"`
		APIKey string `url:"api_key"`
		Num    int    `url:"num"`
		Engine string `url:"engine"`
	}{
		Query:  companyName,
		APIKey: apiKey,
		Num:    1, // Fetch only the first result
		Engine: "google",
	}

	// Encode the query parameters
	queryParams, _ := query.Values(params)
	searchURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())

	// Make the HTTP request to SerpAPI
	resp, err := http.Get(searchURL)
	if err != nil {
		return "", fmt.Errorf("failed to make request to SerpAPI: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	// Decode the JSON response
	var serpResponse SerpAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&serpResponse)
	if err != nil {
		return "", fmt.Errorf("failed to decode SerpAPI response: %w", err)
	}

	// Ensure that we have at least one result
	if len(serpResponse.Organic) == 0 {
		return "", fmt.Errorf("no results found for %s", companyName)
	}

	// Return the first result URL
	return serpResponse.Organic[0].Link, nil
}

func WriteOutput(filename string, output map[string]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for company, email := range output {
		line := fmt.Sprintf("%s: %s\n", company, email)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}
