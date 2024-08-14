package web

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func SearchGoogle(companyName string) (string, error) {
	baseURL := "https://www.google.com/search"
	query := url.Values{}
	query.Add("q", companyName)

	searchUrl := fmt.Sprintf("%s?%s", baseURL, query.Encode())

	// Create a new request
	req, err := http.NewRequest("GET", searchUrl, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set the User-Agent header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to search Google: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 status code: %d", resp.StatusCode)
	}

	// Parse the HTML response
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse Google search results: %w", err)
	}

	// Extract URLs from the search results
	var homepageURL string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists && strings.Contains(link, "/url?q=") {

			fmt.Println("link:", link)

			start := strings.Index(link, "/url?q=") + len("/url?q=")
			end := strings.Index(link[start:], "&")
			if end != -1 {
				homepageURL = link[start : start+end]
			} else {
				homepageURL = link[start:]
			}

			// Filter out social media links
			if !strings.Contains(homepageURL, "facebook.com") && !strings.Contains(homepageURL, "linkedin.com") {
				return
			}
		}
	})

	if homepageURL == "" {
		return "", fmt.Errorf("no homepage URL found")
	}
	return homepageURL, nil
}
