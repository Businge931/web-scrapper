package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func InferAboutPage(homepageURL string) (string, error) {
	resp, err := http.Get(homepageURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch homepage: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse homepage: %w", err)
	}

	var aboutPageURL string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists && (strings.Contains(strings.ToLower(link), "about") || strings.Contains(strings.ToLower(link), "contact")) {
			aboutPageURL = link
			if !strings.HasPrefix(aboutPageURL, "http") {
				aboutPageURL = strings.TrimRight(homepageURL, "/") + "/" + strings.TrimLeft(aboutPageURL, "/")
			}
			return
		}
	})

	if aboutPageURL == "" {
		return "", fmt.Errorf("no About page URL found")
	}

	return aboutPageURL, nil
}
