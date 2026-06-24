package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")
	contsr := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode

	db, err := sql.Open("postgres", contsr)
	if err != nil {
		fmt.Println("error while connecting to the database")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		fmt.Println("database is reachable")
	}

	return db, nil
}
