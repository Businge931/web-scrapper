package companies

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

	var companies []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		companies = append(companies, scanner.Text())
	}
	return companies, scanner.Err()
}