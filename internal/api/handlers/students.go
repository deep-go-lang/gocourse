package handlers

import (
	"fmt"
	"net/http"
)

func StudentsHandler(w http.ResponseWriter, r *http.Request) {
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
