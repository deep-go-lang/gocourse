package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/router"
	"restapi/internal/repository/sqlconnect"

	"github.com/joho/godotenv"
)


func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		return 
	}

	_, err = sqlconnect.ConnectDB()
	if err != nil {
		fmt.Println("Error----", err)
		return
	}


	cert := "cert.pem"
	key := "key.pem"

	port := os.Getenv("API_PORT")

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		}

	// rateLimiter := mw.NewRateLimiter(5, 1*time.Minute)

	// hppMiddleware := mw.Hpp(mw.HPPOptions{
	// 	CheckQuery: true,
	// 	CheckBody: true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist: []string{"name", "age", "city", "class", "sortBy", "sortOrder"},
	// })

	// secureMux := utils.ApplyMiddleware(mux, 
	// 	rateLimiter.Middleware,
	// 	hppMiddleware,
	// 	mw.CORSMiddleware,
	// 	mw.SecurityHeadersMiddleware,
	// 	mw.CompressionMiddleware,
	// 	mw.ResponseTimeMiddleware,
	// )

	secureMux := mw.SecurityHeadersMiddleware(router.Router())

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Printf("Starting server on port %s\n", port)

	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal(err)
	}

}

