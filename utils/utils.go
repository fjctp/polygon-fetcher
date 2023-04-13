package utils

import (
	"log"
	"os"
	"time"
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
func MakeDirIfNotExist(path string) error {
	if !Exist(path) {
		err := os.MkdirAll(path, 0750)
		if err != nil {
			return err
		}
	}
	return nil
}

// Check if a file is modified in the last X amount of years, months, and days
func FileOlderThan(path string, years int, months int, days int) bool {
	if Exist(path) {
		info, _ := os.Stat(path)
		now := time.Now()
		ref := now.AddDate(-years, -months, -days)
		return ref.After(info.ModTime())
	}
	return false
}
