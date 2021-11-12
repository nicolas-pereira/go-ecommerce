package router

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nicolas-pereira/go-ecommerce/server/database"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)
	r.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", http.FileServer((http.Dir("templates")))))
	r.PathPrefix("/styles/").Handler(http.StripPrefix("/styles/", http.FileServer(http.Dir("templates/styles/"))))

	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/product", productHandler)
	r.HandleFunc("/product/{id}", productByIdHandler).Methods("GET")
	return r
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tables, err := database.DatabaseTableCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id, err := database.PostProduct(r.FormValue("name"), r.FormValue("description"), price)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		location := fmt.Sprintf("product/%d", id)
		w.Header().Set("location", location)
		w.WriteHeader(http.StatusSeeOther)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "405 METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
	}
}

func productByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	product, err := database.GetProductById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "404 page not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	tmpl, err := template.ParseFiles("./templates/product.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, product)
}
