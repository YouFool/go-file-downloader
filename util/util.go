package util

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

const CacheFilePath = "./output/_downloadedCache"

// Reads all URLs from a input file into a slice, returns error if could not read file
func ReadURLsFromFilePath(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []string{}, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Error while closing input file: %v", err)
		}
	}()

	return readURLsFromFile(file), nil
}

// Reads URLs from a file which MUST exist
func readURLsFromFile(file *os.File) []string {
	var URLs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		URLs = append(URLs, scanner.Text())
	}
	return URLs
}

// Gets the file name from a given URL
func GetFileNameFromUrl(url string) string {
	tokens := strings.Split(url, "/")
	return tokens[len(tokens)-1]
}

// Creates a new cache file or opens an already existing one and dumps the downloaded URLs into it
func WriteDownloadedURLsToFile(downloadedURLs []string) {
	cacheFile, err := os.OpenFile(CacheFilePath, os.O_CREATE, 0644)

	// close file on exit and check for its returned error
	defer func() {
		if err := cacheFile.Close(); err != nil {
			log.Fatalf("Error while closing output cache file: %v", err)
		}
		log.Println("Done writing cache file!")
	}()

	if errors.Is(err, os.ErrNotExist) {
		// File doesn't exist, adds URLs to the cache file
		updateCacheFileURLs(cacheFile, downloadedURLs)
	} else {
		// Cache already exits, appends URLs to the cache file
		cachedURLs := readURLsFromFile(cacheFile)
		difference := FindDifference(downloadedURLs, cachedURLs)
		updateCacheFileURLs(cacheFile, difference)
	}
}

// Updates the cache file with the new-downloaded URLs
func updateCacheFileURLs(cacheFile *os.File, downloadedURLs []string) {
	for _, u := range downloadedURLs {
		_, err := cacheFile.WriteString(u + "\n")

		if err != nil {
			log.Fatalf("Error while writing URL to cache file: %v", err)
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
