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
	fmt.Println("compnay names",companyNames)
}