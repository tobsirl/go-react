package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB // global variable to hold the connection

func ConnectDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// read in the environment variables
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	// set up postgres sql to open a connection
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("Error connecting to the database: ", errSql)
		panic(errSql)
	} else {
		Db = db
		fmt.Println("Successfully connected to the database")
	}
}
