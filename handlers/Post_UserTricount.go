package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

func UserTricount (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user , err := strconv.Atoi(params["user"])
	checkErr(err)
	tricount, err := strconv.Atoi(params["tricount"])
	checkErr(err)
	var part_user_id int

	rows2, err := db.Query("SELECT COUNT(user_id) AS count FROM Utilisateur WHERE user_id = $1;", user)
	count2 := checkCount(rows2)
	checkErr(err)
	if count2 != 1 {
		fmt.Fprintln(w, "L'utilisateur n'existe pas")
		return
	}

	rows3, err := db.Query("SELECT COUNT(tricount_id) AS count FROM Tricount WHERE tricount_id = $1;", tricount)
	count3 := checkCount(rows3)
	checkErr(err)
	if count3 != 1 {
		fmt.Fprintln(w, "Le Tricount n'existe pas")
		return
	}
	rows, err := db.Query("SELECT COUNT(part_user_id) AS count FROM Participe WHERE part_user_id = $1 AND part_tricount_id = $2;", user,tricount)
	count := checkCount(rows)
	checkErr(err)
	if count == 1 {
		fmt.Fprintln(w, "L'utilisateur appartient déjà à ce Tricount")
		return
	}
	insert := db.QueryRow("INSERT INTO Participe (part_user_id, part_tricount_id, part_balance) SELECT u.user_id, t.tricount_id, 0 FROM Utilisateur u,Tricount t WHERE u.user_id = $1 and t.tricount_id = $2 RETURNING part_user_id;",user,tricount).Scan(&part_user_id)
	checkErr(insert)
	fmt.Fprintln(w, "Ajout de l'utilisateur : ", user, " au Tricount : ", tricount)
}
