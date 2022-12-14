package connections

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewSQLStore() *sql.DB {
	database, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/checkpoint2")
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return database
}
