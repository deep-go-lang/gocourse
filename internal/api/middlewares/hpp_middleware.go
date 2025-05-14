package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

type HPPOptions struct {
	CheckQuery bool
	CheckBody bool
	CheckBodyOnlyForContentType string
	Whitelist []string
}




func Hpp(options HPPOptions) func(http.Handler) http.Handler {
	fmt.Println("Hpp middleware")
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Hpp middleware being returned...")
			if options.CheckBody && r.Method == http.MethodPost && isCorrectContentType(r, options.CheckBodyOnlyForContentType) {
				filterBodyParams(r, options.Whitelist)
			}

			if options.CheckQuery && r.URL.Query() != nil {
				filterQueryParams(r, options.Whitelist)
			}
			next.ServeHTTP(w, r)
			fmt.Println("Hpp middleware serving next handler...")
		})
	}
}

// func isValidHeader(key string, value []string) bool {
// 	for _, v := range value {
// 		if strings.Contains(v, "<") || strings.Contains(v, ">") {
// 			return false
// 		}
// 	}
// 	return true
// }

func isCorrectContentType(r *http.Request, contentType string) bool {
	return strings.Contains(r.Header.Get("Content-Type"), contentType)
}

func filterBodyParams(r *http.Request, whitelist []string) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form", err)
		return
	}
	for k, v := range r.Form {
		if len(v) > 1 {
			// r.Form.Set(k, v[0])
			r.Form.Set(k, v[len(v)-1])
		}

		if !isInWhitelist(k, whitelist) {
			delete(r.Form, k)
		}
	}
}

func isInWhitelist(key string, whitelist []string) bool {
	for _, v := range whitelist {
		if v == key {
			return true
		}
	}
	return false
}

func filterQueryParams(r *http.Request, whitelist []string) {
	query := r.URL.Query()
	for k, v := range query {
		if len(v) > 1 {
			// query.Set(k, v[0])
			query.Set(k, v[len(v)-1])
		}

		if !isInWhitelist(k, whitelist) {
			query.Del(k)
		}
	}

	r.URL.RawQuery = query.Encode()
}
