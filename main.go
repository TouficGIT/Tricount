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
	r.HandleFunc("/getListTricount", handlers.GetListTricount)
	r.HandleFunc("/close", Close)
	log.Fatal(http.ListenAndServe("localhost:8005", r))
	defer db.Close()
}

func Close(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db.Close()
	fmt.Fprintf(w,"Session close !\n")
}

