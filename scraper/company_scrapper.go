package scraper

import (
	"bufio"
	"os"
)

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