package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Reads all URLs from a input file into a slice, returns error if could not read file
func ReadURLsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error while closing input file: %v", err)
		}
	}()

	var URLs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		URLs = append(URLs, scanner.Text())
	}
	return URLs, nil
}

// Gets the file name from a given URL
func GetFileNameFromUrl(url string) string {
	tokens := strings.Split(url, "/")
	return tokens[len(tokens)-1]
}

// Creates an output file and dumps the original URLs
func WriteDownloadedURLsToFile(downloadedURLs []string) {
	file, err := os.Create("./output/_downloadedCache")
	if err != nil {
		log.Fatalf("Error while creating output file: %v", err)
	}

	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error while closing output file: %v", err)
		}
	}()

	for _, u := range downloadedURLs {
		_, err := file.WriteString(u + "\n")

		if err != nil {
			log.Fatalf("Error while writing URL to file: %v", err)
		}
	}
}

// Returns the difference from the elements in `a` that aren't in `b`.
func FindDifference(elements, original []string) []string {
	source := make(map[string]struct{}, len(original))
	for _, item := range original {
		source[item] = struct{}{}
	}
	var diff []string
	for _, elem := range elements {
		if _, found := source[elem]; !found {
			diff = append(diff, elem)
		}
	}
	return diff
}
