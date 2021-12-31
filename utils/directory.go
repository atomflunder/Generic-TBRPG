package utils

import (
	"log"
	"os"
)

//checks if the given directory exists
func DirectoryCheck(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

//creates a new directory
func MakeDirectory(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
