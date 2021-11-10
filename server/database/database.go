package database

import (
	"database/sql"
	"fmt"

	"github.com/nicolas-pereira/go-ecommerce/server/config"

	_ "github.com/go-sql-driver/mysql"
)

var Handler *sql.DB

func InitDb() error {
	var err error
	Handler, err = sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Dbname))
	if err != nil {
		return err
	}
	err = Handler.Ping()
	if err != nil {
		return err
	}
	return err
}
