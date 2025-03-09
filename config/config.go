package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbName, dbPassword, dbSSLMode)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to DB")
	DB = db
	createUserTable()

}

func createUserTable() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
    		id SERIAL PRIMARY KEY,
    		username VARCHAR(255) UNIQUE NOT NULL,
    		firstname VARCHAR(255) NOT NULL,
    		lastname VARCHAR(255) NOT NULL,
    		password TEXT NOT NULL
		);
		`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Error creating users table:", err)
	}
	fmt.Println("Users table is ready!")
}
