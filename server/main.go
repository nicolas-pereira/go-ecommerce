package main

import (
	"log"
	"net/http"

	"github.com/nicolas-pereira/go-ecommerce/server/config"
	"github.com/nicolas-pereira/go-ecommerce/server/database"
	"github.com/nicolas-pereira/go-ecommerce/server/router"
)

func main() {
	err := config.GetConfig()
	if err != nil {
		log.Println(err.Error())
	}
	err = database.InitDb()
	if err != nil {
		log.Println(err.Error())
	}
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
