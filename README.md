# 📁 URL File Downloader

## ❓ Was ist das?
Just a simple `Go` program i've made to download files from a `.txt` file with **a lot** of URL's.

## ⚙ How does it works?
Given a rate limit and a lot of URL's to download, i've had to download files in **parallel and on small batches**, otherwise, the server would deny further newcoming requests with rate limiting.

Since Go has a superb support on [concurrency and parallelism](https://www.youtube.com/watch?v=cN_DpYBzKso) i decided to _give it a go_ (chuckles). Also, i've found good [answers](https://stackoverflow.com/questions/45472324/download-files-by-chunks-in-multiple-threads-in-golang) on StackOverflow that kept me going, many thanks!    