package middlewares

import (
	"fmt"
	"net/http"
)

func SecurityHeadersMiddleware(next http.Handler) http.Handler {
	fmt.Println("SecurityHeadersMiddleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("SecurityHeadersMiddleware being returned...")
		w.Header().Set("X-DNS-Prefetch-Control", "off")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("X-Powered-By", "Django")
		w.Header().Set("Server", "nginx")
		w.Header().Set("X-Permitted-Cross-Domain-Policies", "none")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		
		w.Header().Set("Permissions-Policy", "geolocation=(self), microphone=()")
		next.ServeHTTP(w, r)
		fmt.Println("SecurityHeadersMiddleware serving next handler...")
	})
}
