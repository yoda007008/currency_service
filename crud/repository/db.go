package repository

import (
	"database/sql"
	"fmt"
	"os"
)

func Connect() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")

	dsn := fmt.Sprintf("host=%s port=%s password=%s dbname=%s user=%s sslmode=disable", host, port, password, name, user)
	fmt.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
