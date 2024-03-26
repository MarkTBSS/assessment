package main

import (
	"database/sql"
	"log"

	"github.com/MarkTBSS/assessment/databases"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	//log.Fatal("Error loading .env file: ", godotenv.Load("development.env"))
	err := godotenv.Load("development.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database, err := sql.Open("postgres", databases.DatabaseURL())
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer database.Close()
	createTable := `CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount INT,
		note TEXT,
		tags TEXT[]
	);`
	_, err = database.Exec(createTable)

	if err != nil {
		log.Fatal("can't create table", err)
	}
}
