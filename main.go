package main

import (
	"bufio"
	"fmt"
	"os"
	
)

func readCompanyNames(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var companies []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		companies = append(companies, scanner.Text())
	}
	return companies, scanner.Err()
}



func main() {
	companyNames, err := readCompanyNames("companies.txt")
	if err != nil {
		fmt.Println("Error reading company names:", err)
		return
	}
	fmt.Println(companyNames)
}
