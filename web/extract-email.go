package web

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func ExtractEmailFromAboutPage(aboutPageURL string) (string, error) {
	resp, err := http.Get(aboutPageURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch About page: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read About page: %w", err)
	}

	htmlContent := string(body)
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	email := emailRegex.FindString(htmlContent)

	if email == "" {
		return "", fmt.Errorf("no email address found")
	}

	return email, nil
}