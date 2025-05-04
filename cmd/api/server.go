package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Teacher struct {
	Name []string `json:"name"`
	Age  string   `json:"age"`
	City string   `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, root route!")
	w.Write([]byte("Hello, root route!"))
	fmt.Println("Hello, root route!")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
		
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")
		fmt.Println("userID:", userID)
		fmt.Println("query:", r.URL.Query())

		queryParams := r.URL.Query()
		sortby := queryParams.Get("sortby")
		key := queryParams.Get("key")
		sortorder := queryParams.Get("sortorder")
		category := queryParams.Get("category")
		if sortorder == "" {
			sortorder = "DESC"
		}

		if sortby == "" {
			sortby = "name"
		}

		if key == "" {
			key = "name"
		}

		fmt.Printf("sortby: %s, key: %s, sortorder: %s, category: %s\n", sortby, key, sortorder, category)

		
		w.Write([]byte("GET request teachers route!"))
		fmt.Println("GET request teachers route!")
		return
	case http.MethodPost:
		w.Write([]byte("POST request teachers route!"))
		fmt.Println("POST request teachers route!")
		return
	case http.MethodPut:
		w.Write([]byte("PUT request teachers route!"))
		fmt.Println("PUT request teachers route!")
		return
	case http.MethodDelete:
		w.Write([]byte("DELETE request teachers route!"))
		fmt.Println("DELETE request teachers route!")
		return
	case http.MethodPatch:
		w.Write([]byte("PATCH request teachers route!"))
		fmt.Println("PATCH request teachers route!")
		return
	case http.MethodOptions:
		w.Write([]byte("OPTIONS request teachers route!"))
		fmt.Println("OPTIONS request teachers route!")
		return
	// default:
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	}
	// fmt.Fprintf(w, "Hello, teachers route!")
	w.Write([]byte("Hello, teachers route!"))
	fmt.Println("Hello, teachers route!")
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET request students route!"))
		fmt.Println("GET request students route!")
		return
	case http.MethodPost:
		w.Write([]byte("POST request students route!"))
		fmt.Println("POST request students route!")
		return
	case http.MethodPut:
		w.Write([]byte("PUT request students route!"))
		fmt.Println("PUT request students route!")
		return
	case http.MethodDelete:
		w.Write([]byte("DELETE request students route!"))
		fmt.Println("DELETE request students route!")
		return
	case http.MethodPatch:
		w.Write([]byte("PATCH request students route!"))
		fmt.Println("PATCH request students route!")
		return
	case http.MethodOptions:
		w.Write([]byte("OPTIONS request students route!"))
		fmt.Println("OPTIONS request students route!")
		return
	// default:
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	}
	// fmt.Fprintf(w, "Hello, students route!")
	w.Write([]byte("Hello, students route!"))
	fmt.Println("Hello, students route!")
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET request execs route!"))
		fmt.Println("GET request execs route!")
		return
	case http.MethodPost:
		w.Write([]byte("POST request execs route!"))
		fmt.Println("POST request execs route!")
		return
	case http.MethodPut:
		w.Write([]byte("PUT request execs route!"))
		fmt.Println("PUT request execs route!")
		return
	case http.MethodDelete:
		w.Write([]byte("DELETE request execs route!"))
		fmt.Println("DELETE request execs route!")
		return
	case http.MethodPatch:
		w.Write([]byte("PATCH request execs route!"))
		fmt.Println("PATCH request execs route!")
		return
	case http.MethodOptions:
		w.Write([]byte("OPTIONS request execs route!"))
		fmt.Println("OPTIONS request execs route!")
		return
	// default:
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	}
	
	// fmt.Fprintf(w, "Hello, execs route!")
	w.Write([]byte("Hello, execs route!"))
	fmt.Println("Hello, execs route!")
}

func main() {

	port := 3000

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/teachers/", teachersHandler)

	http.HandleFunc("/students/", studentsHandler)

	http.HandleFunc("/execs/", execsHandler)

	

	fmt.Printf("Starting server on port %d\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}

}
