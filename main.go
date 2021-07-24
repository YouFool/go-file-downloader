package main

import (
	"bytes"
	"file-downloader/util"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

var numFiles int
var numFilesDownloaded int
var downloadedURLs []string

func main() {
	URLs, err := util.ReadURLsFromFile("./input.txt")
	if err != nil {
		log.Fatalf("Error while reading URLs from input file: %v", err)
	}
	downloadedURLs, err = util.ReadURLsFromFile("./output/_downloadedCache")
	if err != nil {
		log.Printf("Could not read URLs from cache file: %v", err)
	} else {
		// TODO: Create slice with remaining images
	}

	numFiles = len(URLs)

	log.Printf("We have %d files to download!", numFiles)

	var wg sync.WaitGroup
	// limit to 50 concurrent downloads
	limiter := make(chan struct{}, 50)
	for _, URL := range URLs {
		wg.Add(1)
		go downloader(&wg, limiter, URL)
	}
	wg.Wait()

	util.WriteDownloadedURLsToFile(downloadedURLs)
}

// Downloads a file using a semaphore to block further requests
func downloader(wg *sync.WaitGroup, semaphore chan struct{}, URL string) {
	semaphore <- struct{}{}
	defer func() {
		<-semaphore
		wg.Done()
	}()

	client := &http.Client{Timeout: 900 * time.Second}
	result, err := client.Get(URL)
	if err != nil {
		log.Fatalf("Error from server while executing request: %v", err)
	}

	defer func() {
		if err := result.Body.Close(); err != nil {
			log.Fatalf("Erroe while reading response body: %v", err)
		}
	}()

	var buf bytes.Buffer
	// I'm copying to a buffer before writing it to file
	// I could also just use IO copy to write it to the file
	// directly and save memory by dumping to the disk directly.
	_, _ = io.Copy(&buf, result.Body)
	// write the bytes to file
	fileName := util.GetFileNameFromUrl(URL)
	_ = ioutil.WriteFile("./output/"+fileName, buf.Bytes(), 0644)
	numFilesDownloaded++
	log.Printf("Downloaded file with name: %s -> (%d/%d)", fileName, numFilesDownloaded, numFiles)
	downloadedURLs = append(downloadedURLs, URL)
	return
}
