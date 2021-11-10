package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicolas-pereira/go-ecommerce/server/config"
	"github.com/nicolas-pereira/go-ecommerce/server/database"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/test", testHandler).Methods("GET")

	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer((staticFileDirectory)))
	r.PathPrefix("static").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/", rootHandler).Methods("GET")
	return r
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var tables int
	err := database.Handler.QueryRow("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA =?", config.Database.Dbname).Scan(&tables)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if tables > 0 {
		http.ServeFile(w, r, "./static/index.html")
	} else {
		fmt.Fprintf(w, "TODO: Database setup\n")
	}
}
