package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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
	query := "SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA =?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelfunc()
	err := handler.QueryRowContext(ctx, query, config.Database().Dbname).Scan(&tables)
	return tables, err
}

// Represents euro amount in cents
type EUR int64

func ToEUR(f float64) EUR {
	return EUR((f * 100) + 0.5)
}

func (e EUR) Float64() float64 {
	return float64(e) / 100
}

func (e EUR) Multiply(f float64) EUR {
	v := (float64(e) * f) + 0.5
	return EUR(v)
}

func (e EUR) String() string {
	f := float64(e)
	f = f / 100
	return fmt.Sprintf("$%.2f", f)
}

// Insert product with given name, description and price, returns inserted product ID
func PostProduct(name string, description string, price float64) (int64, error) {
	query := "INSERT INTO product (name, description, price, fk_category_id) VALUES (?, ?, ?, 1)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancelfunc()

	res, err := handler.ExecContext(ctx, query, name, description, ToEUR(price))
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}
