package main

import (
	"fmt"
)

func runCrawler(localpath string, cloudpath string) {
	hash := getHash(localpath)
	fmt.Printf("%x", hash)
}
