package database

import (
	"database/sql"
	"fmt"

	"github.com/nicolas-pereira/go-ecommerce/server/config"

	_ "github.com/go-sql-driver/mysql"
)

var handler *sql.DB

func init() {
	var err error
	handler, err = sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		config.Database().Username,
		config.Database().Password,
		config.Database().Host,
		config.Database().Port,
		config.Database().Dbname))
	if err != nil {
		handler = nil
	}
}

func Handler() *sql.DB {
	return handler
}

func DatabaseTableCount() (int, error) {
	var tables int
	err := handler.QueryRow("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA =?", config.Database().Dbname).Scan(&tables)
	return tables, err
}
