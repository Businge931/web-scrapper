package main

import (
	"fmt"

	"webscrapper/companies"
	"webscrapper/web"
)

func getCompanyEmail(companyName string) (string, error) {
	homepageURL, err := web.SearchGoogle(companyName)
	if err != nil {
		return "", fmt.Errorf("error searching Google: %w", err)
	}

	fmt.Println("URL:", homepageURL)

	// aboutPageURL, err := web.InferAboutPage(homepageURL)
	// if err != nil {
	// 	return "", fmt.Errorf("error inferring About page: %w", err)
	// }

	// email, err := web.ExtractEmailFromAboutPage(aboutPageURL)
	// if err != nil {
	// 	return "", fmt.Errorf("error extracting email: %w", err)
	// }

	return homepageURL, nil
	// return email, nil
}

func main() {
	_, err := companies.ReadCompanyNames("companies/companies.txt")
	if err != nil {
		fmt.Println("Error reading company names:", err)
		return
	}

	companyName := "stanbic"
	url, err := web.SearchGoogle(companyName)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Company homepage URL found:", url)
	}

	// for _, companyName := range companyNames {
	// 	time.Sleep(2 * time.Second) // Simulate delays to prevent getting blocked by Google

	// 	email, err := getCompanyEmail(companyName)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 	} else {
	// 		fmt.Println("Email found:", email)
	// 	}

	// }

}
