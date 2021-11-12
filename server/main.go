package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/nicolas-pereira/go-ecommerce/server/config"
	"github.com/nicolas-pereira/go-ecommerce/server/database"
	"github.com/nicolas-pereira/go-ecommerce/server/router"
)

func main() {
	if config.Database() == nil {
		log.Println(errors.New("config: can't read database configuration").Error())
	}
	if database.Handler() == nil {
		log.Println(errors.New("database: can't connect to database").Error())
	}
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
