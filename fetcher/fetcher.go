package fetcher

import (
	"errors"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
)

// Get an instance of Fetcher
func NewFetcher() (*polygon.Client, error) {
	key, isFound := os.LookupEnv("POLYGON_API_KEY")
	if isFound && key != "" {
		return polygon.New(key), nil
	} else {
		return nil, errors.New(
			"Environment variable, POLYGON_API_KEY, is empty")
	}
}
