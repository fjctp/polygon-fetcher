package utils

import (
	"log"
	"os"
)

// If there is an err, log it
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Check if a path exists
func Exist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Make a directory if the path does not exist
func MakeDirIfNotExist(path string) {
	if Exist(path) {
		err := os.MkdirAll(path, 0750)
		CheckError(err)
	}
}
