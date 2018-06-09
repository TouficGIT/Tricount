package handlers

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"log"
)

func NewUser (w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	nom := params["nom"]
	var user_id string
	rows, err := db.Query("SELECT COUNT(user_nom) AS count FROM Utilisateur WHERE user_nom = $1;", nom)
	count := checkCount(rows)
	checkErr(err)
	if count == 1 {
		fmt.Fprintln(w, "Nom d'utilisateur déjà existant :", nom)
		fmt.Println("row count : ", count)
	} else {
		insert := db.QueryRow("INSERT INTO Utilisateur (user_nom) VALUES ($1) RETURNING user_id;", nom).Scan(&user_id)

		if insert != nil {
			log.Fatal(insert)
		}

		fmt.Println("Ajout de l'utilisateur : ", nom)
		fmt.Fprintln(w, "Ajout de l'utilisateur : ", nom, " avec l'identifiant : ", user_id)
	}
}






