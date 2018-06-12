package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func DelUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user_nom := params["nom"]
	var user_id int

	rows, err := db.Query("SELECT COUNT(user_id) AS count FROM Utilisateur WHERE user_nom = $1;", user_nom)
	checkErr(err)
	count := checkCount(rows)

	if count != 1 {
		fmt.Fprintln(w, "Utilisateur inexistant :", user_nom)
		return
	}

	row2 :=  db.QueryRow("SELECT user_id FROM Utilisateur WHERE user_nom = $1;",user_nom).Scan(&user_id)
	checkErr(row2)

	fmt.Println(user_id)
	fmt.Println(user_nom)
	del, err := db.Exec("DELETE FROM Participe WHERE part_user_id = $1;",user_id)
	checkErr(err)
	test, err := del.RowsAffected()
	checkErr(err)
	fmt.Println(test)

	del2, err := db.Exec("DELETE FROM Utilisateur WHERE user_id = $1;",user_id)
	checkErr(err)
	test2, err := del2.RowsAffected()
	checkErr(err)
	fmt.Println(test2)

	fmt.Fprintln(w, "Suppression de l'utilisateur : ", user_nom)
}
