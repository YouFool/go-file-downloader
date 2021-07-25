# 📁 URL File Downloader

## ❓ Was ist das?
This is a simple `Go` program that I've made to download files from a `.txt` file with **a lot** of URLs (much more than 10k).
It features a built-in cache, so it will only download un-cached files if an isolated error occurs.

## ⚙ How does it work?
Since the server had a rate limiter, and I've had a lot of URLs to download, it was necessary to download files on **small batches and in parallel**, otherwise, the server would deny further requests with its rate limiter.

Since Go has a superb support to [concurrency and parallelism](https://www.youtube.com/watch?v=cN_DpYBzKso) I decided to _give it a Go_ (chuckles). Also, I've found good [answers](https://stackoverflow.com/questions/45472324/download-files-by-chunks-in-multiple-threads-in-golang) on StackOverflow that kept me going, many thanks!

## 💪 How to I get started?
1. You will need [Go 1.14 or higher](https://golang.org/dl/) installed to compile the source code
2. With Go installed, let's compile the source code
    * Open the terminal and execute `cd path/to/file-downloader && go build`. In example: `cd C:\Users\myUser\Downloads\go-file-downloader && go build` This command will generate an executable file _(.exe)_ inside the folder.
3. With executable generated, create a folder called `output` relative to it
4. Once the program compiles successfully, you just need to provide a file called `input.txt` with the URLs to download, where each link must be separated from the other by a line feed
5. **That's it! You're ready to go!** Run the _.exe_ file and check the downloaded files inside the `output` folder. Just note that:
    * At the end of execution, the downloader will create a cache file called `_downloadedCache` with all downloaded URLs. If some error occurred while downloading file(s) you can just run it again and download the failed ones.
    * You can also delete/rename this cache file to download everything again  

Example program output:
![image](https://user-images.githubusercontent.com/37518972/126886083-709f4488-abdc-411a-bb3e-260158db321b.png)


