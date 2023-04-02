package utils

import "log"

func Check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
