package sqlconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	fmt.Println("Connecting to MariaDB")
	// err := godotenv.Load()
	// if err != nil {
	// 	return nil, err
	// }
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	dbhost := os.Getenv("DB_HOST")
	
	fmt.Printf("Connecting to database with host: %s, port: %s, user: %s, database: %s\n", dbhost, dbport, user, dbName)
	
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, dbhost, dbport, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Printf("Error opening database connection: %v\n", err)
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return nil, err
	}
	fmt.Println("Connected to MariaDB")
	return db, nil
}
