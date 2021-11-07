package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDb() error {
	var err error

	DB, err = sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/",
		os.Getenv("SQL_USER"),
		os.Getenv("SQL_PASSWORD"),
		os.Getenv("SQL_HOST"),
		os.Getenv("SQL_PORT")))
	if err != nil {
		return err
	}
	return DB.Ping()
}
