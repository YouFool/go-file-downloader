package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Reads all URL's from a input file
func ReadUrlsFromInputFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//log.Println(scanner.Text())
		urls = append(urls, scanner.Text())
	}
	return urls
}

// Gets the file name from a given URL
func GetFileNameFromUrl(url string) string {
	tokens := strings.Split(url, "/")
	return tokens[len(tokens)-1]
}
