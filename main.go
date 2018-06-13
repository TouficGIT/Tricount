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

	r.HandleFunc("/joke", handlers.SendJoke).Methods("GET")
	r.HandleFunc("/newTricount/{titre}", handlers.NewTricount).Methods("POST")
	r.HandleFunc("/newUser/{nom}", handlers.NewUser).Methods("POST")
	r.HandleFunc("/setUserTricount/{user}/{tricount}", handlers.UserTricount).Methods("POST")
	r.HandleFunc("/newDepense/{nom}/{montant}/{type}/{titre}", handlers.NewDepense).Methods("POST")
	r.HandleFunc("/refund/{nom}/{tricount}/{montant}",handlers.Refund).Methods("PUT")
	r.HandleFunc("/delUser/{nom}",handlers.DelUser).Methods("DELETE")
	r.HandleFunc("/delTricount/{nom}",handlers.DelTricount).Methods("DELETE")
	r.HandleFunc("/getListTricount", handlers.GetListTricount).Methods("GET")
	r.HandleFunc("/getListUser", handlers.GetListUser).Methods("GET")
	r.HandleFunc("/getListUser2", handlers.GetListUser2).Methods("GET")
	r.HandleFunc("/getDepenseTricount/{titre}", handlers.GetDepenseTricount).Methods("GET")
	r.HandleFunc("/getInfoTricount/{titre}", handlers.GetInfoTricount).Methods("GET")
	r.HandleFunc("/getBalanceTricount/{titre}", handlers.GetBalanceTricount).Methods("GET")

	r.HandleFunc("/close", Close)
	http.Handle("/", r)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	log.Fatal(http.ListenAndServe("localhost:80", r))
	defer db.Close()
}

func Close(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db.Close()
	fmt.Fprintf(w,"Session close !\n")
}

