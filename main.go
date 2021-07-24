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

var numberOfFiles int
var numberOfFilesDownloaded int

func main() {
	urls := ReadUrlsFromInputFile("./input.txt")
	numberOfFiles = len(urls)

	log.Printf("We have %d files to download!", numberOfFiles)

	var wg sync.WaitGroup
	// limit to 50 concurrent downloads
	limiter := make(chan struct{}, 50)
	for _, url := range urls {
		wg.Add(1)
		go downloader(&wg, limiter, url)
	}
	wg.Wait()

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
		log.Fatal(err)
	}
	defer result.Body.Close()
	var buf bytes.Buffer
	// I'm copying to a buffer before writing it to file
	// I could also just use IO copy to write it to the file
	// directly and save memory by dumping to the disk directly.
	_, _ = io.Copy(&buf, result.Body)
	// write the bytes to file
	fileName := util.GetFileNameFromUrl(URL)
	_ = ioutil.WriteFile("./output/"+fileName, buf.Bytes(), 0644)
	numberOfFilesDownloaded++
	log.Printf("Downloaded file with name %s ! (%d/%d)", fileName, numberOfFilesDownloaded, numberOfFiles)
	return
}
