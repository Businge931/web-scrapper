package main

import (
	"fmt"
	"log"

	"github.com/Businge931/company-email-scraper/scraper"
)

func main() {
	companyNames, err := scraper.ReadCompanyNames("data/input.txt")
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}
	output := make(map[string]string)

	for _, companyName := range companyNames {
		companyURL, err := scraper.GetSearchResults(companyName)
		if err != nil {
			log.Printf("Error getting search results for %s: %v", companyName, err)
			output[companyName] = ""
			continue
		}
		fmt.Println("company url", companyURL)
		output[companyName] = companyURL

	}

}
