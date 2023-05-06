package middleware

import (
	"log"
	"net/http"
	"strings"
)

// Check if it is a valid ticker
func IsValidTicker(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

// A middleware that calls f() before calling h.ServeHTTP()
type UpdateData struct {
	h http.Handler
	f Updater
}

// An interface definition for an Updater()
type Updater func(string) error

// ServeHTTP handles the request by passing it to the real
// handler and generate a html report if it does not exist
func (l *UpdateData) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		htmlPage := strings.TrimLeft(r.URL.Path, "/")
		ticker := strings.TrimRight(htmlPage, ".html")
		if ticker != "" && IsValidTicker(ticker) {
			log.Printf("Request report for %s\n", ticker)
			err := l.f(ticker)
			if err != nil {
				// FIXME: return 404
				log.Println(err)
			}
		}
	}
	l.h.ServeHTTP(w, r)
}

func NewUpdateData(h http.Handler, f Updater) *UpdateData {
	return &UpdateData{h, f}
}
