package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ResponseTimeMiddleware(next http.Handler) http.Handler {
	fmt.Println("ResponseTimeMiddleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ResponseTimeMiddleware being returned...")
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)
		rw.Header().Set("X-Response-Time", duration.String())
		
		log.Printf("Request: Method: %s, URL: %s, StatusCode: %d, Duration: %s", r.Method, r.URL.Path, rw.statusCode, duration)
		fmt.Println("ResponseTimeMiddleware serving next handler...")
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}



