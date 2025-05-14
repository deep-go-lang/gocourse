package handlers

import (
	"fmt"
	"net/http"
)

func ExecsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("GET request execs route!"))
		fmt.Println("GET request execs route!")
		return
	case http.MethodPost:
		// fmt.Println("Query:", r.URL.Query())
		// fmt.Println("name:", r.URL.Query().Get("name"))

		// err := r.ParseForm()
		// if err != nil {
		// 	http.Error(w, "Failed to parse form", http.StatusBadRequest)
		// 	return
		// }

		// fmt.Println("Form:", r.Form)
		// fmt.Println("name:", r.Form.Get("name"))
		w.Write([]byte("POST request execs route!"))

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
