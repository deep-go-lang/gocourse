package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func CompressionMiddleware(next http.Handler) http.Handler {
	fmt.Println("CompressionMiddleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("CompressionMiddleware being returned...")
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		w = &gzipWriter{ResponseWriter: w, gz: gz}
		next.ServeHTTP(w, r)

		fmt.Println("CompressionMiddleware serving next handler...")
	})
}

type gzipWriter struct {
	http.ResponseWriter
	gz *gzip.Writer
}

func (w *gzipWriter) Write(b []byte) (int, error) {
	return w.gz.Write(b)
}

func (w *gzipWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}


