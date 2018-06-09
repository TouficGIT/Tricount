package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"log"
)

func NewTricount (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	titre := params["titre"]
	var tricount_id string
	rows, err := db.Query("SELECT COUNT(tricount_titre) AS count FROM Tricount WHERE tricount_titre = $1;", titre)
	count := checkCount(rows)
	checkErr(err)
	if count == 1 {
		fmt.Fprintln(w, "Tricount déjà existant :", titre)
		fmt.Println("row count : ", count)
	} else {
		query := db.QueryRow("INSERT INTO Tricount (tricount_titre) VALUES ($1) RETURNING tricount_id;", titre).Scan(&tricount_id)
		if query != nil {
			log.Fatal(query)
		}
		fmt.Println("Ajout du tricount : ", titre)
		fmt.Fprintln(w, "Ajout du tricount : ", titre, " avec l'identifiant : ", tricount_id)

	}
}
