package main

import "fmt"
import "os"
import "log"
import (
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_ssl_mode := os.Getenv("DB_SSL_MODE")

	connStr := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=%s", db_host, db_user, db_password, db_name, db_port, db_ssl_mode)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT NOW()")

	if err != nil {
		fmt.Println("pq error:", err)
		os.Exit(1)
	}
	_ = rows

	fmt.Println("Connected OK")
	os.Exit(0)
}
