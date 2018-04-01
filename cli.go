package main

import "flag"

func parseCLI() (string, string) {
	localDirectory := flag.String("localpath", "",
		"Path of the local project to monitor")
	cloudPath := flag.String("cloudpath", "",
		"Path to location to monitor on the cloud")
	flag.Parse()

	return *localDirectory, *cloudPath
}

