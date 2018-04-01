package main

func main() {
	localPath, cloudPath := parseCLI()
	// run as go routine?
	runCrawler(localPath, cloudPath)
}


