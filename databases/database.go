package databases

import (
	"database/sql"
	"log"

	"github.com/MarkTBSS/assessment/configs"
	_ "github.com/lib/pq"
)

func ConnectDatabase() error {
	database, err := sql.Open("postgres", configs.DatabaseURL())
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
	return err
}
