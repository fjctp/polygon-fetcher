package middleware

import (
	"log"
	"net/http"
	"time"
)

type HttpLogger struct {
	handler http.Handler
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *HttpLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func NewHttpLogger(h http.Handler) *HttpLogger {
	return &HttpLogger{h}
}
