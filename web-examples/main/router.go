package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	/*	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
		r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
		r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
		r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
	*/
	r.HandleFunc("/books/{title}", BookHandler).Host("localhost:8081")

	/*
		r.HandleFunc("/secure", SecureHandler).Schemes("https")
		r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	*/

	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", AllBooks)
	bookrouter.HandleFunc("/{title}", GetBook)

	http.ListenAndServe(":8080", r)

}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["title"])
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "CreateBook: %v\n", vars["title"])
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ReadBook: %v\n", vars["title"])
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "UpdateBook: %v\n", vars["title"])
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "DeleteBook: %v\n", vars["title"])
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["title"])
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["title"])
}

func SecureHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["title"])
}

func InsecureHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["title"])
}
