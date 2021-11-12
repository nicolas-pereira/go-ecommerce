package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nicolas-pereira/go-ecommerce/server/database"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", http.FileServer((http.Dir("templates")))))
	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("templates/styles/"))))

	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/product", productHandler)
	return r
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tables, err := database.DatabaseTableCount()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if tables > 0 {
		http.ServeFile(w, r, "./templates/index.html")
	} else {
		fmt.Fprintf(w, "TODO: Database setup\n")
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "./templates/productForm.html")
	case "POST":
		r.ParseForm()
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		id, err := database.PostProduct(r.FormValue("name"), r.FormValue("description"), price)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintf(w, "Created product with ID %d\n", id)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "405 METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
	}
}
