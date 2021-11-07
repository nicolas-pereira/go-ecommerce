package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/nicolas-pereira/go-ecommerce/database"
	"github.com/nicolas-pereira/go-ecommerce/router"
)

func main() {
	err := godotenv.Load("prod.env")
	if err != nil {
		fmt.Println("Couldn't load env file")
	}

	err = database.InitDb()
	if err != nil {
		fmt.Println("Can't connect to database")
	}
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
