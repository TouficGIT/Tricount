package main

import (
	_ "github.com/lib/pq"
	"./handlers"
	"fmt"
	"database/sql"
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

var db *sql.DB


func main() {
	r := mux.NewRouter()
	handlers.InitDb()
	r.HandleFunc("/index", handlers.Index)
	r.HandleFunc("/getListTricount", handlers.GetListTricount).Methods("GET")
	r.HandleFunc("/getListUser", handlers.GetListUser).Methods("GET")
	r.HandleFunc("/getDepenseTricount/{id}", handlers.GetDepenseTricount).Methods("GET")
	r.HandleFunc("/getInfoTricount/{id}", handlers.GetInfoTricount).Methods("GET")
	r.HandleFunc("/getBalanceTricount/{id}", handlers.GetBalanceTricount).Methods("GET")
	r.HandleFunc("/close", Close)
	log.Fatal(http.ListenAndServe("localhost:8005", r))
	defer db.Close()
}

func Close(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db.Close()
	fmt.Fprintf(w,"Session close !\n")
}

