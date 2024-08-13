package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
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

func searchGoogle(companyName string) ([]string, error) {
	baseURL := "https://www.google.com/search"
	query := url.Values{}
	query.Add("q", companyName)
	query.Add("num", "10") // Number of results to return (max 10)

	searchUrl := fmt.Sprintf("%s?%s", baseURL, query.Encode())

	// Make the HTTP request
	resp, err := http.Get(searchUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to search Google: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 status code: %d", resp.StatusCode)
	}

	// Parse the HTML response
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Google search results: %w", err)
	}

	// Extract URLs from the search results
	var urls []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")

		if exists && strings.Contains(link, "/url?q=") {
			// fmt.Println("url",link)
			// Extract the actual URL by trimming the "/url?q=" prefix and stopping at the first "&"
			start := strings.Index(link, "/url?q=") + len("/url?q=")
			end := strings.Index(link[start:], "&")
			if end != -1 {
				urls = append(urls, link[start:start+end])
			} else {
				urls = append(urls, link[start:])
			}
		}
	})


	return urls, nil
}

func main() {
	companyNames, err := readCompanyNames("companies.txt")
	if err != nil {
		fmt.Println("Error reading company names:", err)
		return
	}

	for _, companyName := range companyNames {
		// Simulate delays to prevent getting blocked by Google
		time.Sleep(2 * time.Second)

		urls, err := searchGoogle(companyName)
		if err != nil {
			fmt.Println("Error searching Google for", companyName, ":", err)
			continue
		}

		for _, url := range urls {
		fmt.Println(url)
	}
	}

}
