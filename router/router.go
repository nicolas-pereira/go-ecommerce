package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/test", testHandler).Methods("GET")

	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer((staticFileDirectory)))
	r.PathPrefix("static").Handler(staticFileHandler).Methods("GET")

	r.Handle("/", http.FileServer(staticFileDirectory))
	return r
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}
