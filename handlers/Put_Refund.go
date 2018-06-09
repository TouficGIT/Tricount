package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

func Refund (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["nom"]
	tricount := params["tricount"]
	montant , err := strconv.ParseFloat(params["montant"],64)
	checkErr(err)

	rows, err := db.Query("SELECT COUNT(user_id) AS count FROM Utilisateur WHERE user_id = $1;", user)
	checkErr(err)
	count := checkCount(rows)

	if count != 1 {
		fmt.Fprintln(w, "Utilisateur inexistant :", user)
		return
	}

	var part_balance float64
	row5 := db.QueryRow("SELECT part_balance FROM Participe WHERE part_user_id = $1 and part_tricount_id = $2;", user,tricount).Scan(&part_balance)
	checkErr(row5)

	part_balance += montant

	rows6, err := db.Query("UPDATE Participe SET part_balance = $1 WHERE part_user_id = $2 and part_tricount_id = $3;", part_balance, user, tricount)
	checkErr(err)
	checkCount(rows6)

	fmt.Fprintln(w, "L'utilisateur ", user, " a remboursé sa dette sur ", tricount, " de ",montant, " euros")
}