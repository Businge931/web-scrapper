package web

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func SearchGoogle(companyName string) (string, error) {
	baseURL := "https://www.google.com"
	query := url.Values{}
	query.Add("q", companyName)
	// query.Add("num", "10") // Number of results to return (max 10)

	searchUrl := fmt.Sprintf("%s?%s", baseURL, query.Encode())

	// Make the HTTP request
	resp, err := http.Get(searchUrl)
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
			start := strings.Index(link, "/url?q=") + len("/url?q=")
			end := strings.Index(link[start:], "&")
			if end != -1 {
				homepageURL = link[start : start+end]
			} else {
				homepageURL = link[start:]
			}

			if strings.Contains(homepageURL, "facebook.com") || strings.Contains(homepageURL, "linkedin.com") {
				homepageURL = ""
			} else {
				return
			}
		}
	})

	if homepageURL == "" {
		return "", fmt.Errorf("no homepage URL found")
	}
	fmt.Println("url:", homepageURL)
	return homepageURL, nil
}