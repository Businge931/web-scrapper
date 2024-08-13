package main

import (
	"fmt"
	"time"

	"webscrapper/companies"
	"webscrapper/web"
)

func getCompanyEmail(companyName string) (string, error) {
	homepageURL, err := web.SearchGoogle(companyName)
	if err != nil {
		return "", fmt.Errorf("error searching Google: %w", err)
	}

	aboutPageURL, err := web.InferAboutPage(homepageURL)
	if err != nil {
		return "", fmt.Errorf("error inferring About page: %w", err)
	}

	email, err := web.ExtractEmailFromAboutPage(aboutPageURL)
	if err != nil {
		return "", fmt.Errorf("error extracting email: %w", err)
	}

	return email, nil
}

func main() {
	companyNames, err := companies.ReadCompanyNames("companies.txt")
	if err != nil {
		fmt.Println("Error reading company names:", err)
		return
	}

	for _, companyName := range companyNames {
		// Simulate delays to prevent getting blocked by Google
		time.Sleep(2 * time.Second)

		email, err := getCompanyEmail(companyName)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Email found:", email)
		}

	}

}
