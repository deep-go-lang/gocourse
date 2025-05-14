package middlewares

import (
	"fmt"
	"net/http"
)

var allowedOrigins = []string{"https://localhost:3000", "https://localhost:3001", "https://localhost:3002", "https://www.myfrontend.com"}

func CORSMiddleware(next http.Handler) http.Handler {
	fmt.Println("CORSMiddleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("CORSMiddleware being returned...")
		origin := r.Header.Get("Origin")
		fmt.Println("Origin:", origin)
		if isAllowedOrigin(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}else {
			http.Error(w, "Not allowed", http.StatusForbidden)
			return
		}



		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Max-Age", "3600")

		if r.Method == http.MethodOptions {
			return
		}
		// if origin == "" {
		// 	next.ServeHTTP(w, r)
		// 	return
		// }
		// w.Header().Set("Access-Control-Allow-Origin", origin)
		// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		next.ServeHTTP(w, r)
		fmt.Println("CORSMiddleware serving next handler...")
	})
}

func isAllowedOrigin(origin string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if allowedOrigin == origin {
			return true
		}
	}
	return false
}
