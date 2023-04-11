package middleware

import (
	"log"
	"net/http"
	"strings"
)

func IsValidTicker(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

type UpdateData struct {
	h http.Handler
	f Updater
}

type Updater func(string, int, string) error

const num_terms = 20
const term = "Q" // Q: quarterly, A: annually

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *UpdateData) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		htmlPage := strings.TrimLeft(r.URL.Path, "/")
		ticker := strings.TrimRight(htmlPage, ".html")
		if ticker != "" && IsValidTicker(ticker) {
			log.Printf("Request report for %s\n", ticker)
			err := l.f(ticker, num_terms, term)
			if err != nil {
				// FIXME: return 404
			}
		}
	}
	l.h.ServeHTTP(w, r)
}

func NewUpdateData(h http.Handler, f Updater) *UpdateData {
	return &UpdateData{h, f}
}
