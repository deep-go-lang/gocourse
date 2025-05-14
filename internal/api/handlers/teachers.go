package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"strconv"
	"strings"
)

func isValidSortOrder(order string) bool {
	return order == "asc" || order == "desc"
}

func isValidSortField(field string) bool {
	validFields := map[string]bool{
		"id": true,
		"first_name": true,
		"last_name": true,
		"email": true,
		"class": true,
		"subject": true,
	}

	return validFields[field]
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")

	// Common response structure
	type Response struct {
		Status  string      `json:"status"`
		Count   int         `json:"count"`
		Data    interface{} `json:"data"`
		Message string      `json:"message,omitempty"`
	}

	w.Header().Set("Content-Type", "application/json")

	if idStr == "" {

		query := "SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"
		args := []interface{}{}

		query, args = addFilters(r, query, args)

		query, shouldReturn := addSorting(r, query, w)
		if shouldReturn {
			return
		}


		rows, err := db.Query(query, args...)
		if err != nil {
			http.Error(w, "Failed to query teachers", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		teachersList := make([]models.Teacher, 0)
		for rows.Next() {
			var teacher models.Teacher
			err := rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
			if err != nil {
				http.Error(w, "Failed to scan teacher data", http.StatusInternalServerError)
				return
			}
			teachersList = append(teachersList, teacher)
		}

		if err = rows.Err(); err != nil {
			http.Error(w, "Error iterating teacher rows", http.StatusInternalServerError)
			return
		}

		response := Response{
			Status: "success",
			Count:  len(teachersList),
			Data:   teachersList,
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	// Handle single teacher by ID
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response := Response{
			Status:  "error",
			Message: "Invalid ID format",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Query single teacher
	var teacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).
		Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
	
	if err != nil {
		if err == sql.ErrNoRows {
			response := Response{
				Status:  "error",
				Message: "Teacher not found",
			}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}
		http.Error(w, "Failed to query teacher", http.StatusInternalServerError)
		return
	}

	// Return single teacher
	response := Response{
		Status: "success",
		Count:  1,
		Data:   teacher,
	}
	json.NewEncoder(w).Encode(response)
}

func addSorting(r *http.Request, query string, w http.ResponseWriter) (string, bool) {
	sortParams := r.URL.Query()["sortby"]

	if len(sortParams) > 0 {
		query += " ORDER BY"
		for i, param := range sortParams {
			parts := strings.Split(param, ":")
			if len(parts) != 2 {
				http.Error(w, "Invalid sort parameter", http.StatusBadRequest)
				return "", true
			}
			field := parts[0]
			order := parts[1]
			if !isValidSortField(field) || !isValidSortOrder(order) {
				continue
			}
			if i > 0 {
				query += ","
			}

			query += " " + field + " " + order
		}
	}
	return query, false
}

func addFilters(r *http.Request, query string, args []interface{}) (string, []interface{}) {
	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"email":      "email",
		"class":      "class",
		"subject":    "subject",
	}

	for param, dbField := range params {
		value := r.URL.Query().Get(param)
		if value != "" {
			query += " AND " + dbField + " = ?"
			args = append(args, value)
		}
	}
	return query, args
}

func createTeacherHandler(w http.ResponseWriter, r *http.Request) {
	// mutex.Lock()
	// defer mutex.Unlock()



	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}

	defer db.Close()
	

	// Common response structure
	type Response struct {
		Status  string      `json:"status"`
		Count   int         `json:"count"`
		Data    interface{} `json:"data"`
		Message string      `json:"message,omitempty"`
	}

	w.Header().Set("Content-Type", "application/json")

	var newTeachers []models.Teacher
	err = json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		response := Response{
			Status:  "error",
			Message: "Invalid request body",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	addedTeachers := make([]models.Teacher, 0, len(newTeachers))

	for _, newTeacher := range newTeachers {
		resp, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to execute statement: %v", err), http.StatusInternalServerError)
			return
		}
		id, err := resp.LastInsertId()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get last insert id: %v", err), http.StatusInternalServerError)
			return
		}
		newTeacher.ID = int(id)
		addedTeachers = append(addedTeachers, newTeacher)
	}
	

	

	// Return success response
	response := Response{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}



func TeachersHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)
	switch r.Method {
	case http.MethodGet:
		getTeachersHandler(w, r)
	case http.MethodPost:
		createTeacherHandler(w, r)
	case http.MethodPut:
		//PUT Method
		updateTeacherHandler(w, r)
	case http.MethodDelete:
		w.Write([]byte("DELETE request teachers route!"))
		fmt.Println("DELETE request teachers route!")
		return
	case http.MethodPatch:
		patchTeacherHandler(w, r)
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

//PUT Method
func updateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/teachers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedTeacher models.Teacher
	err = json.NewDecoder(r.Body).Decode(&updatedTeacher)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).
		Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to query teacher", http.StatusInternalServerError)
		return
	}

	updatedTeacher.ID = existingTeacher.ID

	stmt, err := db.Prepare("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatedTeacher.FirstName, updatedTeacher.LastName, updatedTeacher.Email, updatedTeacher.Class, updatedTeacher.Subject, updatedTeacher.ID)
	if err != nil {
		http.Error(w, "Failed to execute statement", http.StatusInternalServerError)
		return
	}

	type Response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
	}

	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Status: "success",
		Message: "Teacher updated successfully",
	}
	json.NewEncoder(w).Encode(response)

	

}

func patchTeacherHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/teachers/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	


	db, err := sqlconnect.ConnectDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).
		Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)
	
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to query teacher", http.StatusInternalServerError)
		return
	}

	for key, value := range updates {
		switch key {
		case "first_name":
			existingTeacher.FirstName = value.(string)
		case "last_name":
			existingTeacher.LastName = value.(string)
		case "email":
			existingTeacher.Email = value.(string)
		case "class":
			existingTeacher.Class = value.(string)
		case "subject":
			existingTeacher.Subject = value.(string)
		}
	}

	stmt, err := db.Prepare("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(existingTeacher.FirstName, existingTeacher.LastName, existingTeacher.Email, existingTeacher.Class, existingTeacher.Subject, existingTeacher.ID)
	if err != nil {
		http.Error(w, "Failed to execute statement", http.StatusInternalServerError)
		return
	}

	type Response struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
	}

	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Status: "success",
		Message: "Teacher updated successfully",
	}
	json.NewEncoder(w).Encode(response)
	
	
}
