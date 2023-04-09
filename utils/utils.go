package utils

import (
	"log"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func MakeDir(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0750)
		CheckError(err)
	}
}
