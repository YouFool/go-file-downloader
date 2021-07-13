# 📁 URL File Downloader

## ❓ Was ist das?
This is a simple `Go` program that I've made to download files from a `.txt` file with **a lot** of URL's (much more than 10k).

## ⚙ How does it work?
Since the server had a rate limiter, and I've had a lot of URL's to download, it was necessary to download files on **small batches and in parallel **, otherwise, the server would deny further requests with its rate limiter.

Since Go has a superb support to [concurrency and parallelism](https://www.youtube.com/watch?v=cN_DpYBzKso) I decided to _give it a Go_ (chuckles). Also, I've found good [answers](https://stackoverflow.com/questions/45472324/download-files-by-chunks-in-multiple-threads-in-golang) on StackOverflow that kept me going, many thanks!    