package main

import (
	"flag"
	"fmt"
)

func main() {
	localPath, cloudPath := parseCLI()
	fmt.Printf("Local Project Path: %s\n", localPath)
	fmt.Printf("Cloud Path: %s\n", cloudPath)
}

func parseCLI() (string, string) {
	localDirectory := flag.String("localpath", "",
		"Path of the local project to monitor")
	cloudPath := flag.String("cloudpath", "",
		"Path to location to monitor on the cloud")
	flag.Parse()

	return *localDirectory, *cloudPath
}
