package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func UserTricount (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := (params["user"])
	tricount := (params["tricount"])
	var part_user_id int
	var user_id int
	var tricount_id int

	rows2, err := db.Query("SELECT COUNT(user_id) AS count FROM Utilisateur WHERE user_nom = $1;", user)
	count2 := checkCount(rows2)
	checkErr(err)
	if count2 != 1 {
		fmt.Fprintln(w, "L'utilisateur n'existe pas")
		return
	}

	rows3, err := db.Query("SELECT COUNT(tricount_id) AS count FROM Tricount WHERE tricount_titre = $1;", tricount)
	count3 := checkCount(rows3)
	checkErr(err)
	if count3 != 1 {
		fmt.Fprintln(w, "Le Tricount n'existe pas")
		return
	}

	query := db.QueryRow ("SELECT user_id FROM Utilisateur WHERE user_nom = $1;", user).Scan(&user_id)
	checkErr(query)

	query2 := db.QueryRow ("SELECT tricount_id FROM Tricount WHERE tricount_titre = $1;", tricount).Scan(&tricount_id)
	checkErr(query2)

	rows, err := db.Query("SELECT COUNT(part_user_id) AS count FROM Participe WHERE part_user_id = $1 AND part_tricount_id = $2;", user_id,tricount_id)
	count := checkCount(rows)
	checkErr(err)
	if count == 1 {
		fmt.Fprintln(w, "L'utilisateur appartient déjà à ce Tricount")
		return
	}
	insert := db.QueryRow("INSERT INTO Participe (part_user_id, part_tricount_id, part_balance) SELECT u.user_id, t.tricount_id, 0 FROM Utilisateur u,Tricount t WHERE u.user_id = $1 and t.tricount_id = $2 RETURNING part_user_id;",user_id,tricount_id).Scan(&part_user_id)
	checkErr(insert)
	fmt.Fprintln(w, "Ajout de l'utilisateur : ", user, " au Tricount : ", tricount)
}
