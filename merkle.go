package main

import (
	"crypto/sha256"
	"io/ioutil"
)

// todo: build the merkle hash tree on the file

func getHash(path string) (hash []byte) {
	data, err := ioutil.ReadFile(path)
	check(err)
	hasher := sha256.New()
	_, err = hasher.Write(data)
	check(err)
	hash = hasher.Sum(nil)
	return
}

func check(e error) {
	if e!= nil {
		panic(e)
	}
}
